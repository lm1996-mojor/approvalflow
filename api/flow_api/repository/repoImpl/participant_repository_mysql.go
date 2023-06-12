package repoImpl

import (
	"five.com/lk_flow/api/common"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type ParticipantRepositoryImpl struct {
}

func ParticipantNewMysqlRepository() repository.ParticipantRepository {
	return ParticipantRepositoryImpl{}
}

func (m ParticipantRepositoryImpl) BatchInsertParticipantInfo(tx *gorm.DB, pointValueId int64, participantInfos []model.Participant) (pointParticipantList []model.Participant) {
	for i := 0; i < len(participantInfos); i++ {
		var participant model.Participant
		trans.CopyFields(participantInfos[i], &participant)
		participant.PointValueId = pointValueId
		participant.ApprovalResult = 5
		if err := tx.Table(common.Participant.TableName()).Create(&participant).Error; err != nil {
			log.Error("新增参与者信息出错")
			panic("服务器错误")
		}
		participantInfos[i].Id = participant.Id
		participantInfos[i].ApprovalResult = 5
		participantInfos[i].PointValueId = pointValueId

	}
	return participantInfos
}

// SelectParticipantListByPointValueId 根据节点值id和指定条件查询参与人列表
func (m ParticipantRepositoryImpl) SelectParticipantListByPointValueId(db *gorm.DB, pointValueId int64, conditionMap map[string]interface{}) (participantList []model.Participant) {
	db = db.Table(common.Participant.TableName())
	if conditionMap != nil {
		for k, v := range conditionMap {
			db = db.Where(k, v)
		}
	}
	if pointValueId != 0 {
		db = db.Where("point_value_id = ?", pointValueId)
	}
	db.Select(common.Participant.GetAllColumn()).Scan(&participantList)
	return participantList
}

// UpdateParticipantInfo 更新参与人数据
func (m ParticipantRepositoryImpl) UpdateParticipantInfo(tx *gorm.DB, participantList []model.Participant) bool {
	for _, participant := range participantList {
		if err := tx.Table(common.Participant.TableName()).Updates(participant).Error; err != nil {
			log.Error("更新参与人信息错误: " + err.Error())
			return false
		}
	}
	return true
}

// DeleteParticipantInfo 删除参与人数据
func (m ParticipantRepositoryImpl) DeleteParticipantInfo(tx *gorm.DB, participantList []model.Participant) bool {
	for _, detail := range participantList {
		if err := tx.Table(common.Participant.TableName()).Where("id = ?", detail.Id).Delete(&common.Participant).Error; err != nil {
			log.Error("删除参与人信息错误: " + err.Error())
			return false
		}
	}
	return true
}
