package control_info

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/utils"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainCtlList(mdl ListQueryMdl) rest.Result       //根据条件查询控件列表
	ObtainSingleCtlInfo(id int64) rest.Result         //查询单个控件信息
	AddCtl(infos []model.ControlInfo) rest.Result     //新增控件
	ChangeCtlById(info model.ControlInfo) rest.Result //修改控件
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainCtlList 根据条件查询控件列表
func (s ServiceImpl) ObtainCtlList(mdl ListQueryMdl) rest.Result {
	db := lib.ObtainCustomDb()
	ctlList := s.repo.SelectCtlList(db, mdl)
	resultMap := make(map[string]interface{}, 1)
	resultMap["ctlList"] = ctlList
	return rest.SuccessResult(resultMap)
}

// ObtainSingleCtlInfo 查询单个控件信息
func (s ServiceImpl) ObtainSingleCtlInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	ctlInfo := s.repo.SelectSingleCtlInfo(db, id)
	if reflect.DeepEqual(ctlInfo, CtlDetail{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	resultMap := make(map[string]interface{})
	resultMap["ctlInfo"] = ctlInfo
	return rest.SuccessResult(resultMap)
}

// AddCtl 新增控件
func (s ServiceImpl) AddCtl(infos []model.ControlInfo) rest.Result {
	tx := lib.ObtainCustomDbTx()
	for i := 0; i < len(infos); i++ {
		infos[i].CtlCode = utils.GenerateCodeByUUID(3)
	}
	s.repo.InsertCtl(tx, infos)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// ChangeCtlById 修改控件
func (s ServiceImpl) ChangeCtlById(info model.ControlInfo) rest.Result {
	db := lib.ObtainCustomDb()
	ctlInfo := s.repo.SelectSingleCtlInfo(db, info.Id)
	//以防编码被改动(保持原本的编码)
	info.CtlCode = ctlInfo.CtlCode
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateCtl(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}
