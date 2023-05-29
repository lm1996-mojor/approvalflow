package process

import (
	"reflect"

	"five.com/lk_flow/api/process_group"
	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
	"five.com/technical_center/core_library.git/utils/trans"
)

// Service 业务逻辑接口
type Service interface {
	ObtainProcessPage(mdl QueryListMdl) rest.Result //根据条件获取流程列表
	ObtainProcessInfo(id int64) rest.Result         //获取单个流程信息
	AddProcess(info model.Process) rest.Result      //新增流程
	ChangeProcess(info model.Process) rest.Result   //根据id修改流程信息
	DropProcess(ids []int64) rest.Result            //删除流程
	ReleaseProcess(id int64) rest.Result            //发布流程
}

type ServiceImpl struct {
	repo      Repository
	groupRepo process_group.Repository
}

func NewService(repository Repository, groupRepo process_group.Repository) Service {
	return &ServiceImpl{repo: repository, groupRepo: groupRepo}
}

// ObtainProcessPage 根据条件获取流程列表
func (s ServiceImpl) ObtainProcessPage(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	ids := make([]int64, 0)
	ids[0] = mdl.GroupId
	groupList := s.groupRepo.SelectGroupListByIds(db, ids)
	processList := make([]ListProcessDetail, len(groupList))
	trans.DeepCopy(groupList, processList)
	processList = s.repo.SelectProcessListByCondition(db, processList, mdl)
	return utils.Result(processList, "processList")
}

// ObtainProcessInfo 获取单个流程信息
func (s ServiceImpl) ObtainProcessInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	processInfo := s.repo.SelectProcessInfo(db, id)
	if reflect.DeepEqual(processInfo, DetailProcess{}) {
		log.Error("查询单个流程信息失败")
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(processInfo, "processInfo")
}

// AddProcess 新增流程
func (s ServiceImpl) AddProcess(info model.Process) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertProcess(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// ChangeProcess 根据id修改流程信息
func (s ServiceImpl) ChangeProcess(info model.Process) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateProcess(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}

// DropProcess 删除流程
func (s ServiceImpl) DropProcess(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteProcess(tx, ids)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}

func (s ServiceImpl) ReleaseProcess(id int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	columns := make(map[string]interface{}, 1)
	columns["process_status"] = 1
	s.repo.UpdateProcessSpecifyColumns(tx, id, columns)
	return rest.SuccessCustom("发布流程成功", nil, rest.Success)
}
