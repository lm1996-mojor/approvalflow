package repoImpl

import (
	"five.com/lk_flow/api/common"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type ProcessValueRepositoryImpl struct {
}

func ProcessValueNewMysqlRepository() repository.ProcessValueRepository {
	return ProcessValueRepositoryImpl{}
}

var pointValueRepo = PointValueNewMysqlRepository()
var ctlValueRepo = CtlValueNewMysqlRepository()

func (m ProcessValueRepositoryImpl) InsertProcessValueInfo(tx *gorm.DB, params api_model.ApprovalParams) (approvalParams api_model.ApprovalParams) {
	processValueInfo := getSingleProcessValueInfoObj(params)
	if err := tx.Table(common.ProcessValue.TableName()).Create(&processValueInfo).Error; err != nil {
		log.Error("新增流程值信息出错")
		panic("服务器错误")
	}
	params.Id = processValueInfo.Id
	// 配置当前所执行的审批信息
	// 新增流程节点
	params.PointDetails = pointValueRepo.BatchInsertPointValueInfos(tx, params.PointDetails)
	// 新增流程表单控件值信息
	ctlValueRepo.BatchInsertCtlValueInfos(tx, params.CtlValues)
	return params
}

// SelectProcessValueInfoByCondition 根据条件查询流程值信息
func (m ProcessValueRepositoryImpl) SelectProcessValueInfoByCondition(db *gorm.DB, condition map[string]interface{}) (processValueInfo model.ProcessValue) {
	db = db.Table(common.ProcessValue.TableName())
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Select(common.ProcessValue.GetAllColumn()).Scan(&processValueInfo)
	return processValueInfo
}

// UpdateProcessValueInfoByApprovalCode 根据审批编号更新流程值信息
func (m ProcessValueRepositoryImpl) UpdateProcessValueInfoByApprovalCode(tx *gorm.DB, params api_model.ApprovalParams) bool {
	processValueInfo := getSingleProcessValueInfoObj(params)
	resultStatus := pointValueRepo.UpdatePointValueInfo(tx, params.PointDetails)
	// 更新节点信息失败
	if !resultStatus {
		return false
	}
	if err := tx.Table(common.Participant.TableName()).Updates(processValueInfo).Error; err != nil {
		log.Error("更新流程信息错误: " + err.Error())
		return false
	}
	tx.Commit()
	return true
}

// DeleteProcessValueInfo 删除流程值数据
func (m ProcessValueRepositoryImpl) DeleteProcessValueInfo(tx *gorm.DB, params api_model.ApprovalParams) bool {
	resultStatus := pointValueRepo.DeletePointValueInfo(tx, params.PointDetails)
	if !resultStatus {
		return false
	}
	if err := tx.Table(common.Participant.TableName()).Where("id = ?", params.Id).Delete(common.ProcessValue).Error; err != nil {
		log.Error("删除流程信息错误: " + err.Error())
		return false
	}
	tx.Commit()
	return true
}

// 获取单个流程值对象
func getSingleProcessValueInfoObj(params api_model.ApprovalParams) (processValueInfo model.ProcessValue) {
	trans.CopyFields(params, &processValueInfo)
	return processValueInfo
}
