package condition_detail_groups

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectGroupListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectGroupListByIds(db *gorm.DB, ids []int64) (groupList []model.ConditionDetailGroups) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ConditionDetailGroups.TableName()).Select(ConditionDetailGroups.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageConditionDetailGroupsByCondition 根据条件获取条件详细信息分组列表
func (m MysqlRepository) SelectPageConditionDetailGroupsByCondition(db *gorm.DB, mdl QueryListMdl) (items []GroupsDetail) {
	var groupsInfos []model.ConditionDetailGroups
	db.Table(ConditionDetailGroups.TableName()).Where("condition_info_id = ?", mdl.ConditionInfoId).Select(ConditionDetailGroups.GetAllColumn()).Scan(&groupsInfos)
	trans.DeepCopy(groupsInfos, &items)
	for i := 0; i < len(groupsInfos); i++ {
		var conditionDetailInfos []model.ConditionDetailInfo
		db.Table(ConditionDetailInfo.TableName()).Where("groups_id = ?").Select(ConditionDetailInfo.GetAllColumn()).Scan(&conditionDetailInfos)
		items[i].ConditionDetailInfos = conditionDetailInfos
	}
	return
}

// SelectConditionDetailGroupsById 获取单个条件详细信息分组的信息
func (m MysqlRepository) SelectConditionDetailGroupsById(db *gorm.DB, id int64) (processGroupsInfo model.ConditionDetailGroups) {
	db.Table(ConditionDetailGroups.TableName()).Where("id = ?", id).Select(ConditionDetailGroups.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertConditionDetailGroups 新增条件详细信息分组
func (m MysqlRepository) InsertConditionDetailGroups(tx *gorm.DB, info model.ConditionDetailGroups) {
	if err := tx.Table(ConditionDetailGroups.TableName()).Create(&info).Error; err != nil {
		log.Error("新增条件详细信息分组信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateConditionDetailGroupsById 根据id修改条件详细信息分组信息
func (m MysqlRepository) UpdateConditionDetailGroupsById(tx *gorm.DB, info model.ConditionDetailGroups) {
	if err := tx.Table(ConditionDetailGroups.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改条件详细信息分组信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteConditionDetailGroupsById 删除条件详细信息分组
func (m MysqlRepository) DeleteConditionDetailGroupsById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ConditionDetailGroups.TableName()).Delete(&ConditionDetailGroups, ids).Error; err != nil {
		log.Error("删除条件详细信息分组信息出错")
		panic("删除出错，数据已回滚")
	}
}
