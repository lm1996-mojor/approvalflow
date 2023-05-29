package process_group

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainProcessGroupsList(mdl QueryPageMdl) rest.Result         //根据条件获取流程分组列表
	ObtainProcessGroupsInfo(id int64) rest.Result                 //获取单个流程分组的信息
	AddProcessGroups(info model.ProcessGroups) rest.Result        //新增流程分组
	ChangeProcessGroups(info model.ProcessGroups) rest.Result     //根据id修改流程分组信息
	DropProcessGroups(ids []int64) rest.Result                    //删除流程分组
	ObtainProcessGroupsAllList(queryMdl QueryListMdl) rest.Result //获取全部流程分组列表
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainProcessGroupsList 根据条件获取流程分组列表
func (s ServiceImpl) ObtainProcessGroupsList(mdl QueryPageMdl) rest.Result {
	db := lib.ObtainCustomDb()

	//字符串转数字
	mdl.ClientId = utils.GetClientId()
	items, total := s.repo.SelectPageProcessGroupsByCondition(db, mdl)
	return rest.NewQueryPage(items, mdl.PageNumber, mdl.PageSize, total)
}

// ObtainProcessGroupsInfo 获取单个流程分组的信息
func (s ServiceImpl) ObtainProcessGroupsInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	var processGroupsInfo model.ProcessGroups
	processGroupsInfo = s.repo.SelectProcessGroupsById(db, id)
	if reflect.DeepEqual(processGroupsInfo, model.ProcessGroups{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(processGroupsInfo, "processGroupsInfo")
}

// AddProcessGroups 新增流程分组
func (s ServiceImpl) AddProcessGroups(info model.ProcessGroups) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertProcessGroups(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeProcessGroups 根据id修改流程分组信息
func (s ServiceImpl) ChangeProcessGroups(info model.ProcessGroups) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateProcessGroupsById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropProcessGroups 删除流程分组
func (s ServiceImpl) DropProcessGroups(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteProcessGroupsById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}

// ObtainProcessGroupsAllList 获取全部流程分组列表
func (s ServiceImpl) ObtainProcessGroupsAllList(mdl QueryListMdl) rest.Result {
	//字符串转数字
	mdl.ClientId = utils.GetClientId()
	db := lib.ObtainCustomDb()
	processGroupsDetail := s.repo.SelectProcessGroupsList(db, mdl)
	return utils.Result(processGroupsDetail, "processGroupsDetail")
}
