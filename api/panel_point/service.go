package panel_point

import (
	"five.com/lk_flow/api/process"
	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainPanelPointList(mdl QueryListMdl) rest.Result  //根据条件获取流程节点列表
	ObtainPanelPointInfo(id int64) rest.Result          //获取单个流程节点信息
	AddPanelPoint(info PanelPointDetail) rest.Result    //新增流程节点
	ChangePanelPoint(info model.PanelPoint) rest.Result //根据id修改流程节点信息
	DropPanelPoint(ids []int64) rest.Result             //删除流程节点
}

type ServiceImpl struct {
	repo        Repository
	processRepo process.Repository
}

func NewService(repository Repository, processRepo process.Repository) Service {
	return &ServiceImpl{repo: repository, processRepo: processRepo}
}

// ObtainPanelPointList 根据条件获取流程节点列表
func (s ServiceImpl) ObtainPanelPointList(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	panelPointList := s.repo.SelectPanelPointListByCondition(db, map[string]interface{}{"process_id = ?": mdl.ProcessId})
	return utils.Result(panelPointList, "panelPointList")
}

// ObtainPanelPointInfo 获取单个流程节点信息
func (s ServiceImpl) ObtainPanelPointInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	panelPointInfo := s.repo.SelectSinglePanelPointInfoById(db, id)
	return utils.Result(panelPointInfo, "panelPointInfo")
}

// AddPanelPoint 新增流程节点
func (s ServiceImpl) AddPanelPoint(info PanelPointDetail) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertPanelPoint(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// ChangePanelPoint 根据id修改流程节点信息
func (s ServiceImpl) ChangePanelPoint(info model.PanelPoint) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdatePanelPoint(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}

// DropPanelPoint 删除流程节点
func (s ServiceImpl) DropPanelPoint(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeletePanelPoint(tx, ids)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}
