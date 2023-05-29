package condition_detail_info

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectGroupListByIds(db *gorm.DB, ids []int64) []model.ConditionDetailInfo                                  //获取指定的分组信息
	SelectPageConditionDetailInfoByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ConditionDetailInfo) //根据条件获取条件详细信息分组列表
	SelectConditionDetailInfoById(db *gorm.DB, id int64) model.ConditionDetailInfo                              //获取单个条件详细信息分组的信息
	InsertConditionDetailInfo(tx *gorm.DB, info model.ConditionDetailInfo)                                      //新增条件详细信息分组
	UpdateConditionDetailInfoById(tx *gorm.DB, info model.ConditionDetailInfo)                                  //根据id修改条件详细信息分组信息
	DeleteConditionDetailInfoById(tx *gorm.DB, ids []int64)                                                     //删除条件详细信息分组
	UpdateGroupsInfo(tx *gorm.DB, groupId int64)                                                                //修改指定分组信息
}
