package participant_format

import (
	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainParticipantFormatPage(mdl QueryListMdl) rest.Result         //根据条件获取参与者形式列表
	ObtainParticipantFormatInfo(id int64) rest.Result                 //获取单个参与者形式信息
	AddParticipantFormat(info model.ParticipantFormat) rest.Result    //新增参与者形式
	ChangeParticipantFormat(info model.ParticipantFormat) rest.Result //根据id修改参与者形式信息
	DropParticipantFormat(ids []int64) rest.Result                    //删除参与者形式
	ObtainParticipantFormatAllList() rest.Result                      //获取全部参与者形式
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainParticipantFormatPage 根据条件获取参与者形式分页
func (s ServiceImpl) ObtainParticipantFormatPage(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items, total := s.repo.SelectParticipantFormatListByCondition(db, mdl)
	return rest.NewQueryPage(items, mdl.PageNumber, mdl.PageSize, total)
}

// ObtainParticipantFormatInfo 获取单个参与者形式信息
func (s ServiceImpl) ObtainParticipantFormatInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	participantFormatInfo := s.repo.SelectParticipantFormatInfo(db, id)
	return utils.Result(participantFormatInfo, "participantFormatInfo")
}

// AddParticipantFormat 新增参与者形式
func (s ServiceImpl) AddParticipantFormat(info model.ParticipantFormat) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertParticipantFormat(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.Success)
}

// ChangeParticipantFormat 根据id修改参与者形式信息
func (s ServiceImpl) ChangeParticipantFormat(info model.ParticipantFormat) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateParticipantFormat(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}

// DropParticipantFormat 删除参与者形式
func (s ServiceImpl) DropParticipantFormat(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteParticipantFormat(tx, ids)
	return rest.SuccessCustom("修改成功", nil, rest.Success)
}

// ObtainParticipantFormatAllList 获取全部参与者形式
func (s ServiceImpl) ObtainParticipantFormatAllList() rest.Result {
	db := lib.ObtainCustomDb()
	participantFormatList := s.repo.SelectParticipantFormatAllList(db)
	return utils.Result(participantFormatList, "participantFormatList")
}
