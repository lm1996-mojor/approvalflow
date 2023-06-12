package repoImpl

import (
	. "five.com/lk_flow/api/common"
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
	if err := tx.Table(ProcessValue.TableName()).Create(&processValueInfo).Error; err != nil {
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
	db = db.Table(ProcessValue.TableName())
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Select(ProcessValue.GetAllColumn()).Scan(&processValueInfo)
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
	if err := tx.Table(ProcessValue.TableName()).Updates(processValueInfo).Error; err != nil {
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
	if err := tx.Table(ProcessValue.TableName()).Where("id = ?", params.Id).Delete(&ProcessValue).Error; err != nil {
		log.Error("删除流程信息错误: " + err.Error())
		return false
	}
	tx.Commit()
	return true
}

// SelectApprovalOfUserParticipation 分页查询用户参与的审批
func (m ProcessValueRepositoryImpl) SelectApprovalOfUserParticipation(db *gorm.DB, params api_model.QueryApprovalParam) (items []api_model.ApprovalParams, total int64) {
	//AppCode      string `json:"appCode"`      // 应用编码 -- 必填
	//BusinessCode string `json:"businessCode"` // 业务编码
	//ClientId     int64  `json:"-"`            // 租户id
	//ProcessId    int64  `json:"processId"`    // 流程id
	//UserId       int64  `json:"userId"`       // 用户id(即参与人id)
	//QueryType    int8   `json:"qt"`           // 查询类型（1 待处理 2 已处理 3 已发起 4 收到的）
	var processValueList []model.ProcessValue
	db = db.Table(ProcessValue.TableName() + " prv")
	db = db.Joins("LEFT JOIN " + PointValue.TableName() + " pov on pov.approval_code = prv.approval_code")
	db = db.Joins("LEFT JOIN " + Participant.TableName() + " pp on pp.point_value_id = pov.id")
	db = db.Where("prv.app_code = ?", params.AppCode)
	db = db.Where("prv.client_id = ?", params.ClientId)
	db = db.Where("prv.deleted_at is null")
	db = db.Where("prv.process_rate != ?", 6)
	if params.UserId != 0 {
		db = db.Where("pp.obj_id = ?", params)
	}
	if params.ProcessId != 0 {
		db = db.Where("prv.process_id = ?", params.ProcessId)
	}
	db = db.Count(&total)
	params.PageNumber = params.PageNumber * params.PageSize
	db = db.Select("DISTINCT " + ProcessValue.GetAllColumWithAlias("prv"))
	db.Limit(params.PageSize).Offset(params.PageNumber).Scan(&processValueList)
	trans.DeepCopy(processValueList, &items)
	return
}

// 获取单个流程值对象
func getSingleProcessValueInfoObj(params api_model.ApprovalParams) (processValueInfo model.ProcessValue) {
	trans.CopyFields(params, &processValueInfo)
	return processValueInfo
}

// SelectSingleProcessValue 根据id获取该审批信息
func (m ProcessValueRepositoryImpl) SelectSingleProcessValue(db *gorm.DB, approvalCode string) (approvalParams api_model.ApprovalParams) {
	var processValueInfo model.ProcessValue
	if err := db.Table(ProcessValue.TableName()).Where("approval_code = ?", approvalCode).Select(ProcessValue.GetAllColumn()).Scan(&processValueInfo).Error; err != nil {
		log.Error("查询审批信息失败: " + err.Error())
		panic("服务器错误")
	}
	trans.DeepCopy(processValueInfo, &approvalParams)
	approvalParams.PointDetails = pointValueRepo.SelectPointValueInfoByApprovalCode(db, approvalCode)
	approvalParams.CtlValues = ctlValueRepo.SelectCtlValueListByApprovalCode(db, approvalCode)
	return approvalParams
}

// SelectInitiatedApprovalByUserId 根据用户id查询已发起的审批流程
//
// isFilterRevocationStatus：是否过滤流程进度为撤销状态的数据
func (m ProcessValueRepositoryImpl) SelectInitiatedApprovalByUserId(db *gorm.DB, params api_model.QueryApprovalParam) (items []api_model.ApprovalParams, total int64) {
	var processValueList []model.ProcessValue
	// 表连接配置
	db = db.Table(ProcessValue.TableName() + " prv")
	db = db.Joins("LEFT JOIN " + PointValue.TableName() + " pov on pov.approval_code = prv.approval_code")
	db = db.Joins("LEFT JOIN " + PanelPoint.TableName() + " po on po.id = pov.point_id")
	db = db.Joins("LEFT JOIN " + Participant.TableName() + " pp on pp.point_value_id = pov.id")
	// 固定条件
	db = db.Where("prv.app_code = ?", params.AppCode)
	db = db.Where("prv.client_id = ?", params.ClientId)
	db = db.Where("po.point_type = ?", 5)
	db = db.Where("prv.deleted_at is null")
	// 可选条件
	if params.UserId != 0 {
		db = db.Where("pp.obj_id = ?", params.UserId)
	}
	if params.ProcessId != 0 {
		db = db.Where("prv.process_id = ?", params.ProcessId)
	}
	// 结果
	db = db.Count(&total)
	params.PageNumber = params.PageNumber * params.PageSize
	db = db.Select("DISTINCT " + ProcessValue.GetAllColumWithAlias("prv"))
	db.Limit(params.PageSize).Offset(params.PageNumber).Scan(&processValueList)
	trans.DeepCopy(processValueList, &items)
	return
}

// SelectProcessValueListByApprovalList 根据审批编号切片查询审批流程分页列表
func (m ProcessValueRepositoryImpl) SelectProcessValueListByApprovalList(db *gorm.DB, approvalCodeList []string, params api_model.QueryApprovalParam) (items []api_model.ApprovalParams, total int64) {
	var processValueList []model.ProcessValue
	db = db.Table(ProcessValue.TableName())
	db = db.Where("app_code = ?", params.AppCode)
	db = db.Where("client_id = ?", params.ClientId)
	db = db.Where("approval_code in (?)", approvalCodeList)
	db = db.Where("prv.deleted_at is null")
	if params.ProcessId != 0 {
		db = db.Where("process_id = ?", params.ProcessId)
	}
	db = db.Count(&total)
	params.PageNumber = params.PageNumber * params.PageSize
	db = db.Select(ProcessValue.GetAllColumn())
	db.Limit(params.PageSize).Offset(params.PageNumber).Scan(&processValueList)
	trans.DeepCopy(processValueList, &items)
	return
}
