package repository

import (
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// ProcessValueRepository 操作数据库接口
type ProcessValueRepository interface {
	InsertProcessValueInfo(tx *gorm.DB, params api_model.ApprovalParams) (approvalParams api_model.ApprovalParams) // 新增流程值信息
	SelectProcessValueInfoByCondition(db *gorm.DB, condition map[string]interface{}) model.ProcessValue            // 根据条件查询单个流程信息
	UpdateProcessValueInfoByApprovalCode(tx *gorm.DB, params api_model.ApprovalParams) bool                        // 更新流程值数据
	DeleteProcessValueInfo(db *gorm.DB, params api_model.ApprovalParams) bool                                      // 删除流程值数据
}
