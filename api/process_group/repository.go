package process_group

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectGroupListByIds(db *gorm.DB, ids []int64) []model.ProcessGroups                             //获取指定的分组信息
	SelectPageProcessGroupsByCondition(db *gorm.DB, mdl QueryPageMdl) ([]model.ProcessGroups, int64) //根据条件获取流程分组列表
	SelectProcessGroupsById(db *gorm.DB, id int64) model.ProcessGroups                               //获取单个流程分组的信息
	InsertProcessGroups(tx *gorm.DB, info model.ProcessGroups)                                       //新增流程分组
	UpdateProcessGroupsById(tx *gorm.DB, info model.ProcessGroups)                                   //根据id修改流程分组信息
	DeleteProcessGroupsById(tx *gorm.DB, ids []int64)                                                //删除流程分组
	SelectProcessGroupsList(db *gorm.DB, mdl QueryListMdl) []ProcessGroupsDetail                     //获取全部流程分组列表
}
