package condition_detail_info

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectGroupListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectGroupListByIds(db *gorm.DB, ids []int64) (groupList []model.ConditionDetailInfo) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ConditionDetailInfo.TableName()).Select(ConditionDetailInfo.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageConditionDetailInfoByCondition 根据条件获取条件详细信息分组列表
func (m MysqlRepository) SelectPageConditionDetailInfoByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ConditionDetailInfo) {
	db = db.Table(ConditionDetailInfo.TableName())
	db = db.Where("groups_id = ?", mdl.GroupsId)
	db.Select(ConditionDetailInfo.GetAllColumn()).Scan(&items)
	return
}

// SelectConditionDetailInfoById 获取单个条件详细信息分组的信息
func (m MysqlRepository) SelectConditionDetailInfoById(db *gorm.DB, id int64) (processGroupsInfo model.ConditionDetailInfo) {
	db.Table(ConditionDetailInfo.TableName()).Where("id = ?", id).Select(ConditionDetailInfo.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertConditionDetailInfo 新增条件详细信息分组
func (m MysqlRepository) InsertConditionDetailInfo(tx *gorm.DB, info model.ConditionDetailInfo) {
	if err := tx.Table(ConditionDetailInfo.TableName()).Create(&info).Error; err != nil {
		log.Error("新增条件详细信息分组信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateConditionDetailInfoById 根据id修改条件详细信息分组信息
func (m MysqlRepository) UpdateConditionDetailInfoById(tx *gorm.DB, info model.ConditionDetailInfo) {
	if err := tx.Table(ConditionDetailInfo.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改条件详细信息分组信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteConditionDetailInfoById 删除条件详细信息分组
func (m MysqlRepository) DeleteConditionDetailInfoById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ConditionDetailInfo.TableName()).Delete(&ConditionDetailInfo, ids).Error; err != nil {
		log.Error("删除条件详细信息分组信息出错")
		panic("删除出错，数据已回滚")
	}
}

// UpdateGroupsInfo 修改指定分组信息
func (m MysqlRepository) UpdateGroupsInfo(tx *gorm.DB, groupId int64) {
	if err := tx.Table(ConditionDetailInfo.TableName()).Where("id = ?", groupId).Updates(map[string]interface{}{"pass_ask": 1}).Error; err != nil {
		log.Error("修改条件详细信息分组信息出错")
		panic("修改出错，数据已回滚")
	}
}
