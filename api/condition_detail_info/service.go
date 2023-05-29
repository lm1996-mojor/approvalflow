package condition_detail_info

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainConditionDetailInfoList(mdl QueryListMdl) rest.Result           //根据条件获取条件详细信息列表
	ObtainConditionDetailInfoInfo(id int64) rest.Result                   //获取单个条件详细信息的信息
	AddConditionDetailInfo(info model.ConditionDetailInfo) rest.Result    //新增条件详细信息
	ChangeConditionDetailInfo(info model.ConditionDetailInfo) rest.Result //根据id修改条件详细信息信息
	DropConditionDetailInfo(ids []int64) rest.Result                      //删除条件详细信息
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainConditionDetailInfoList 根据条件获取条件详细信息列表
func (s ServiceImpl) ObtainConditionDetailInfoList(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items := s.repo.SelectPageConditionDetailInfoByCondition(db, mdl)
	return utils.Result(items, "conditionDetailInfos")
}

// ObtainConditionDetailInfoInfo 获取单个条件详细信息的信息
func (s ServiceImpl) ObtainConditionDetailInfoInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	conditionDetailInfo := s.repo.SelectConditionDetailInfoById(db, id)
	if reflect.DeepEqual(conditionDetailInfo, model.ConditionDetailInfo{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(conditionDetailInfo, "conditionDetailInfo")
}

// AddConditionDetailInfo 新增条件详细信息
func (s ServiceImpl) AddConditionDetailInfo(info model.ConditionDetailInfo) rest.Result {
	// 查看分组中是否已有条件
	var queryMdl QueryListMdl
	queryMdl.GroupsId = info.GroupsId
	items := s.repo.SelectPageConditionDetailInfoByCondition(lib.ObtainCustomDb(), queryMdl)
	if len(items) > 0 {
		s.repo.UpdateGroupsInfo(lib.ObtainCustomDbTx(), info.GroupsId)
	}
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertConditionDetailInfo(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeConditionDetailInfo 根据id修改条件详细信息信息
func (s ServiceImpl) ChangeConditionDetailInfo(info model.ConditionDetailInfo) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateConditionDetailInfoById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropConditionDetailInfo 删除条件详细信息
func (s ServiceImpl) DropConditionDetailInfo(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteConditionDetailInfoById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}
