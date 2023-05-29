package approval_rule_go

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectPageCtlTabByCondition(db *gorm.DB, mdl QueryPageMdl) ([]model.CtlTab, int64) //根据条件获取条件要素（控件标签）列表
	SelectCtlTabById(db *gorm.DB, id int64) model.CtlTab                               //获取单个条件要素（控件标签）的信息
	InsertCtlTab(tx *gorm.DB, info model.CtlTab)                                       //新增条件要素（控件标签）
	UpdateCtlTabById(tx *gorm.DB, info model.CtlTab)                                   //根据id修改条件要素（控件标签）信息
	DeleteCtlTabById(tx *gorm.DB, ids []int64)                                         //删除条件要素（控件标签）
	SelectCtlTabAllList(db *gorm.DB) (ctlTabList []model.CtlTab)                       //获取所有条件要素（控件标签）
}
