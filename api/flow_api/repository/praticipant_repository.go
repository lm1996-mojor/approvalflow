package repository

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// ParticipantRepository 操作数据库接口
type ParticipantRepository interface {
	BatchInsertParticipantInfo(tx *gorm.DB, pointValueId int64, participantInfos []model.Participant) (pointParticipantList []model.Participant)    // 批量新增参与人信息
	SelectParticipantListByPointValueId(db *gorm.DB, pointValueId int64, conditionMap map[string]interface{}) (participantList []model.Participant) // 根据节点值id和指定条件查询参与人列表
	UpdateParticipantInfo(tx *gorm.DB, participantList []model.Participant) bool                                                                    // 更新参与人数据
	DeleteParticipantInfo(tx *gorm.DB, participantList []model.Participant) bool                                                                    // 删除参与人数据
}
