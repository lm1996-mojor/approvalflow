package repository

import (
	"five.com/lk_flow/api/flow_api/api_model"
	"gorm.io/gorm"
)

// CtlValueRepository 操作数据库接口
type CtlValueRepository interface {
	BatchInsertCtlValueInfos(tx *gorm.DB, ctlValueInfos []api_model.CtlDetail)               // 批量新增控件值信息
	SelectCtlValueListByApprovalCode(db *gorm.DB, approvalCode string) []api_model.CtlDetail // 根据审批编号获取指定的审批表单
}
