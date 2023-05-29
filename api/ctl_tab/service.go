package approval_rule_go

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainCtlTabPage(mdl QueryPageMdl) rest.Result //根据条件获取条件要素（控件标签）列表
	ObtainCtlTabInfo(id int64) rest.Result         //获取单个条件要素（控件标签）的信息
	AddCtlTab(info model.CtlTab) rest.Result       //新增条件要素（控件标签）
	ChangeCtlTab(info model.CtlTab) rest.Result    //根据id修改条件要素（控件标签）信息
	DropCtlTab(ids []int64) rest.Result            //删除条件要素（控件标签）
	ObtainCtlTabAllList() rest.Result              //获取所有条件要素（控件标签）
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainCtlTabPage 根据条件获取条件要素（控件标签）列表
func (s ServiceImpl) ObtainCtlTabPage(mdl QueryPageMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items, total := s.repo.SelectPageCtlTabByCondition(db, mdl)
	return rest.NewQueryPage(items, mdl.PageNumber, mdl.PageSize, total)
}

// ObtainCtlTabInfo 获取单个条件要素（控件标签）的信息
func (s ServiceImpl) ObtainCtlTabInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	ctlTabInfo := s.repo.SelectCtlTabById(db, id)
	if reflect.DeepEqual(ctlTabInfo, model.CtlTab{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(ctlTabInfo, "ctlTabInfo")
}

// AddCtlTab 新增条件要素（控件标签）
func (s ServiceImpl) AddCtlTab(info model.CtlTab) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertCtlTab(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeCtlTab 根据id修改条件要素（控件标签）信息
func (s ServiceImpl) ChangeCtlTab(info model.CtlTab) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateCtlTabById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropCtlTab 删除条件要素（控件标签）
func (s ServiceImpl) DropCtlTab(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteCtlTabById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}

// ObtainCtlTabAllList 获取所有条件要素（控件标签）
func (s ServiceImpl) ObtainCtlTabAllList() rest.Result {
	db := lib.ObtainCustomDb()
	ctlTabList := s.repo.SelectCtlTabAllList(db)
	return utils.Result(ctlTabList, "ctlTabList")
}
