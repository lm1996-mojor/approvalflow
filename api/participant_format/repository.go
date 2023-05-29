package participant_format

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectParticipantFormatListByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ParticipantFormat, total int64) //根据条件获取参与者形式列表
	SelectParticipantFormatInfo(db *gorm.DB, id int64) model.ParticipantFormat                                           //获取单个参与者形式信息
	InsertParticipantFormat(tx *gorm.DB, info model.ParticipantFormat) int64                                             //新增参与者形式
	UpdateParticipantFormat(tx *gorm.DB, info model.ParticipantFormat)                                                   //根据id修改参与者形式信息
	DeleteParticipantFormat(tx *gorm.DB, ids []int64)                                                                    //删除参与者形式
	SelectParticipantFormatAllList(db *gorm.DB) []model.ParticipantFormat                                                //获取所有参与者形式
}
