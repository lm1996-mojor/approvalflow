package panel_point

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectPanelPointListByCondition(db *gorm.DB, condition map[string]interface{}) []PanelPointDetail                      // 根据条件获取流程节点列表
	SelectSinglePanelPointInfoById(db *gorm.DB, id int64) model.PanelPoint                                                 // 获取单个流程节点信息
	InsertPanelPoint(tx *gorm.DB, info PanelPointDetail) int64                                                             // 新增流程节点
	UpdatePanelPoint(tx *gorm.DB, info model.PanelPoint)                                                                   // 根据id修改流程节点信息
	DeletePanelPoint(tx *gorm.DB, ids []int64)                                                                             // 删除流程节点
	SelectSpecifyNodesNextStepInfoByPointId(db *gorm.DB, pointId int64) (panelPointList []model.PanelPoint)                // 根据指定的节点id查询其下一个节点的信息
	SelectSinglePanelPointInfoByCondition(db *gorm.DB, condition map[string]interface{}) (panelPointList PanelPointDetail) // 根据条件获取单个节点及其关系描述信息
}
