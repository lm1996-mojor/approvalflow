package participant_format

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

// SelectParticipantFormatListByCondition 根据条件获取参与者形式列表
func (m MysqlRepository) SelectParticipantFormatListByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ParticipantFormat, total int64) {
	db = db.Table(ParticipantFormat.TableName())
	if mdl.Search != "" {
		db = db.Where("format_name like ?", "%"+mdl.Search+"%")
	}
	db = db.Count(&total)
	db = db.Order("created_at desc")
	mdl.PageNumber = mdl.PageNumber * mdl.PageSize
	if mdl.PageSize <= 0 {
		mdl.PageSize = 10
	}
	db = db.Select(ParticipantFormat.GetAllColumn())
	db.Limit(mdl.PageSize).Offset(mdl.PageNumber).Scan(&items)
	return items, total
}

// SelectParticipantFormatInfo 获取单个参与者形式信息
func (m MysqlRepository) SelectParticipantFormatInfo(db *gorm.DB, id int64) (processInfo model.ParticipantFormat) {
	db.Table(ParticipantFormat.TableName()).Where("id = ?", id).Select(ParticipantFormat.GetAllColumn()).Scan(&processInfo)
	return processInfo
}

// InsertParticipantFormat 新增参与者形式
func (m MysqlRepository) InsertParticipantFormat(tx *gorm.DB, info model.ParticipantFormat) int64 {
	if err := tx.Table(ParticipantFormat.TableName()).Create(&info).Error; err != nil {
		log.Error("新增参与者形式信息出错")
		panic("服务器错误")
	}
	return info.Id
}

// UpdateParticipantFormat 根据id修改参与者形式信息
func (m MysqlRepository) UpdateParticipantFormat(tx *gorm.DB, info model.ParticipantFormat) {
	if err := tx.Table(ParticipantFormat.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改参与者形式信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteParticipantFormat 删除参与者形式
func (m MysqlRepository) DeleteParticipantFormat(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ParticipantFormat.TableName()).Delete(&ParticipantFormat, ids).Error; err != nil {
		log.Error("删除参与者形式信息出错")
		panic("删除出错，数据已回滚")
	}
}

// SelectParticipantFormatAllList 获取所有参与者形式
func (m MysqlRepository) SelectParticipantFormatAllList(db *gorm.DB) (participantFormatList []model.ParticipantFormat) {
	db.Table(ParticipantFormat.TableName()).Select(ParticipantFormat.GetAllColumn()).Scan(&participantFormatList)
	return
}
