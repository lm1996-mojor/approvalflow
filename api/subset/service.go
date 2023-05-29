package subset

import (
	"reflect"

	"five.com/lk_flow/model"

	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/utils"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainSubList(mdl ListQueryMdl) rest.Result //根据条件查询子集列表（分页）
	ObtainSingleSubInfo(id int64) rest.Result   //查询单个子集信息
	AddSub(info model.Subset) rest.Result       //新增子集
	ChangeSubById(info SubDetail) rest.Result   //修改子集
	DropSub(ids []int64) rest.Result            //删除子集
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainSubList 根据条件查询子集列表（分页）
func (s ServiceImpl) ObtainSubList(mdl ListQueryMdl) rest.Result {
	db := lib.ObtainCustomDb()
	subList := s.repo.SelectSubList(db, mdl)
	resultMap := make(map[string]interface{}, 1)
	resultMap["subList"] = subList
	return rest.SuccessResult(resultMap)
}

// ObtainSingleSubInfo 查询单个子集信息
func (s ServiceImpl) ObtainSingleSubInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	subInfo := s.repo.SelectSubInfo(db, id)
	if reflect.DeepEqual(subInfo, SubDetail{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	resultMap := make(map[string]interface{})
	resultMap["subInfo"] = subInfo
	return rest.SuccessResult(resultMap)
}

// AddSub 新增子集
func (s ServiceImpl) AddSub(info model.Subset) rest.Result {
	if info.ParentId > 0 {
		params := make(map[string]interface{})
		params["have_children"] = 1
		s.repo.UpdateSubSpecifyColumnsById(lib.ObtainCustomDbTx(), info.ParentId, params)
	}
	tx := lib.ObtainCustomDbTx()
	info.SubCode = utils.GenerateCodeByUUID(4)
	s.repo.InsertSub(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// ChangeSubById 修改子集
func (s ServiceImpl) ChangeSubById(info SubDetail) rest.Result {
	if info.ParentId > 0 {
		subInfo := s.repo.SelectSubInfo(lib.ObtainCustomDb(), info.Id)
		var subsetInfos []model.Subset
		subsetInfos = s.repo.SelectSubsetInfosByParentId(lib.ObtainCustomDb(), subInfo.ParentId)
		if len(subsetInfos) <= 1 {
			params := make(map[string]interface{})
			params["have_children"] = 2
			s.repo.UpdateSubSpecifyColumnsById(lib.ObtainCustomDbTx(), info.ParentId, params)
		}
	}
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateSub(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// DropSub 删除子集
func (s ServiceImpl) DropSub(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteSub(tx, ids)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}
