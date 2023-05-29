package control_info

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

// SelectCtlList 根据条件查询控件列表
func (m MysqlRepository) SelectCtlList(db *gorm.DB, mdl ListQueryMdl) (items []CtlDetail) {
	if mdl.OwnerType <= 0 {
		mdl.OwnerType = 1
	}
	db.Table(ControlInfo.TableName()).Where("parent_id = ?", mdl.ParentId).Where("owner_type = ?", mdl.OwnerType).Scan(&items)
	return items
}

// SelectSingleCtlInfo 查询单个控件信息
func (m MysqlRepository) SelectSingleCtlInfo(db *gorm.DB, id int64) (ctlInfo CtlDetail) {
	db.Table(ControlInfo.TableName()+" ic").Where("ic.id = ?", id).Select(ControlInfo.GetAllColumWithAlias("ic")).Scan(&ctlInfo)
	return ctlInfo
}

// InsertCtl 新增控件
func (m MysqlRepository) InsertCtl(tx *gorm.DB, info []model.ControlInfo) {
	if err := tx.Unscoped().Table(ControlInfo.TableName()).Where("parent_id = ?", info[0].ParentId).Delete(&ControlInfo).Error; err != nil {
		log.Error("删除控件出错")
		panic("删除出错，数据已回滚")
	}
	if err := tx.Table(ControlInfo.TableName()).CreateInBatches(&info, 20).Error; err != nil {
		log.Error("新增控件出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateCtl 修改控件
func (m MysqlRepository) UpdateCtl(tx *gorm.DB, info model.ControlInfo) {
	if err := tx.Table(ControlInfo.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改控件出错")
		panic("修改出错，数据已回滚")
	}
}
