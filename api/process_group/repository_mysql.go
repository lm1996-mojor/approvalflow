package process_group

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	lib "five.com/technical_center/core_library.git/utils/repo"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectGroupListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectGroupListByIds(db *gorm.DB, ids []int64) (groupList []model.ProcessGroups) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ProcessGroups.TableName()).Select(ProcessGroups.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageProcessGroupsByCondition 根据条件获取流程分组列表
func (m MysqlRepository) SelectPageProcessGroupsByCondition(db *gorm.DB, mdl QueryPageMdl) (items []model.ProcessGroups, total int64) {
	db = db.Table(ProcessGroups.TableName())
	if mdl.Search != "" || len(mdl.Search) > 0 {
		db = db.Where("group_name like ?", "%"+mdl.Search+"%")
	}
	db = db.Where("client_id = ?", mdl.ClientId)
	db = db.Where("app_code = ?", mdl.AppCode)
	db = db.Count(&total)
	db = db.Order("order_no asc")
	mdl.PageNumber = mdl.PageNumber * mdl.PageSize
	if mdl.PageSize <= 0 {
		mdl.PageSize = 10
	}
	db.Select(ProcessGroups.GetAllColumn())
	db.Limit(mdl.PageSize).Offset(mdl.PageNumber).Scan(&items)
	return
}

// SelectProcessGroupsById 获取单个流程分组的信息
func (m MysqlRepository) SelectProcessGroupsById(db *gorm.DB, id int64) (processGroupsInfo model.ProcessGroups) {
	db.Table(ProcessGroups.TableName()).Where("id = ?", id).Select(ProcessGroups.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertProcessGroups 新增流程分组
func (m MysqlRepository) InsertProcessGroups(tx *gorm.DB, info model.ProcessGroups) {
	if err := tx.Table(ProcessGroups.TableName()).Create(&info).Error; err != nil {
		log.Error("新增流程分组信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateProcessGroupsById 根据id修改流程分组信息
func (m MysqlRepository) UpdateProcessGroupsById(tx *gorm.DB, info model.ProcessGroups) {
	if err := tx.Table(ProcessGroups.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改流程分组信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteProcessGroupsById 删除流程分组
func (m MysqlRepository) DeleteProcessGroupsById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ProcessGroups.TableName()).Delete(&ProcessGroups, ids).Error; err != nil {
		log.Error("删除流程分组信息出错")
		panic("删除出错，数据已回滚")
	}
}

func (m MysqlRepository) SelectProcessGroupsList(db *gorm.DB, mdl QueryListMdl) (processGroupsDetail []ProcessGroupsDetail) {
	var processGroupsInfoList []model.ProcessGroups
	db = db.Table(ProcessGroups.TableName())
	db = db.Where("client_id = ?", mdl.ClientId)
	db = db.Where("app_code = ?", mdl.AppCode)
	db.Select(ProcessGroups.GetAllColumn()).Scan(&processGroupsInfoList)
	trans.DeepCopy(processGroupsInfoList, &processGroupsDetail)
	for i := 0; i < len(processGroupsDetail); i++ {
		var processList []model.Process
		lib.ObtainCustomDb().Table(Process.TableName()).Where("group_id = ?", processGroupsDetail[i].Id).Select(Process.GetAllColumn()).Scan(&processList)
		processGroupsDetail[i].ProcessList = processList
	}
	return
}
