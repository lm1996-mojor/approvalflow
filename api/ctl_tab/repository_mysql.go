package approval_rule_go

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

// SelectPageCtlTabByCondition 根据条件获取条件要素（控件标签）列表
func (m MysqlRepository) SelectPageCtlTabByCondition(db *gorm.DB, mdl QueryPageMdl) (items []model.CtlTab, total int64) {
	db = db.Table(CtlTab.TableName())
	if mdl.Search != "" || len(mdl.Search) > 0 {
		db = db.Where("rule_name like ?", "%"+mdl.Search+"%")
	}
	db = db.Count(&total)
	db = db.Order("created_at asc")
	mdl.PageNumber = mdl.PageNumber * mdl.PageSize
	if mdl.PageSize <= 0 {
		mdl.PageSize = 10
	}
	db.Select(CtlTab.GetAllColumn())
	db.Limit(mdl.PageSize).Offset(mdl.PageNumber).Scan(&items)
	return
}

// SelectCtlTabById 获取单个条件要素（控件标签）的信息
func (m MysqlRepository) SelectCtlTabById(db *gorm.DB, id int64) (processGroupsInfo model.CtlTab) {
	db.Table(CtlTab.TableName()).Where("id = ?", id).Select(CtlTab.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertCtlTab 新增条件要素（控件标签）
func (m MysqlRepository) InsertCtlTab(tx *gorm.DB, info model.CtlTab) {
	if err := tx.Table(CtlTab.TableName()).Create(&info).Error; err != nil {
		log.Error("新增条件要素（控件标签）信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateCtlTabById 根据id修改条件要素（控件标签）信息
func (m MysqlRepository) UpdateCtlTabById(tx *gorm.DB, info model.CtlTab) {
	if err := tx.Table(CtlTab.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改条件要素（控件标签）信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteCtlTabById 删除条件要素（控件标签）
func (m MysqlRepository) DeleteCtlTabById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(CtlTab.TableName()).Delete(&CtlTab, ids).Error; err != nil {
		log.Error("删除条件要素（控件标签）信息出错")
		panic("删除出错，数据已回滚")
	}
}

// SelectCtlTabAllList 获取所有条件要素（控件标签）
func (m MysqlRepository) SelectCtlTabAllList(db *gorm.DB) (ctlTabList []model.CtlTab) {
	db.Table(CtlTab.TableName()).Select(CtlTab.GetAllColumn()).Scan(&ctlTabList)
	return
}
