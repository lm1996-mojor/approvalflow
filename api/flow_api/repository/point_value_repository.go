package repository

import (
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// PointValueRepository 操作数据库接口
type PointValueRepository interface {
	BatchInsertPointValueInfos(tx *gorm.DB, pointDetails []api_model.PointDetail) (nodeList []api_model.PointDetail)                // 批量新增流程节点值信息
	SelectPointValueListByCondition(db *gorm.DB, conditionMap map[string]interface{}) (pointValueList []model.PointValue)           // 根据条件查询节点值列表
	SelectPointValueInfoByPointIdAndApprovalCode(db *gorm.DB, pointId int64, approvalCode string) (pointValueInfo model.PointValue) // 根据节点id和审批编码查询指定的节点值信息
	UpdatePointValueInfo(tx *gorm.DB, details []api_model.PointDetail) bool                                                         // 更新节点值数据
	DeletePointValueInfo(tx *gorm.DB, details []api_model.PointDetail) bool                                                         // 删除节点值数据
	SelectPointValueInfoByApprovalCode(db *gorm.DB, approvalCode string) (nodeList []api_model.PointDetail)                         // 根据审批编号查询节点列表
}
