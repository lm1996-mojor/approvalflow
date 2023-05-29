package condition_detail_groups

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectGroupListByIds(db *gorm.DB, ids []int64) []model.ConditionDetailGroups                     //获取指定的分组信息
	SelectPageConditionDetailGroupsByCondition(db *gorm.DB, mdl QueryListMdl) (items []GroupsDetail) //根据条件获取条件详细信息分组列表
	SelectConditionDetailGroupsById(db *gorm.DB, id int64) model.ConditionDetailGroups               //获取单个条件详细信息分组的信息
	InsertConditionDetailGroups(tx *gorm.DB, info model.ConditionDetailGroups)                       //新增条件详细信息分组
	UpdateConditionDetailGroupsById(tx *gorm.DB, info model.ConditionDetailGroups)                   //根据id修改条件详细信息分组信息
	DeleteConditionDetailGroupsById(tx *gorm.DB, ids []int64)                                        //删除条件详细信息分组
}
