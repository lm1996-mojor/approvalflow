package repoImpl

import (
	"five.com/lk_flow/api/common"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/api/panel_point"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	lib "five.com/technical_center/core_library.git/utils/repo"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type PointValueRepositoryImpl struct {
}

func PointValueNewMysqlRepository() repository.PointValueRepository {
	return PointValueRepositoryImpl{}
}

var participantRepo = ParticipantNewMysqlRepository()
var pointRepo = panel_point.NewMysqlRepository()

// BatchInsertPointValueInfos 批量新增流程节点值信息
func (m PointValueRepositoryImpl) BatchInsertPointValueInfos(tx *gorm.DB, pointDetails []api_model.PointDetail) (pointDetailList []api_model.PointDetail) {
	for i := 0; i < len(pointDetails); i++ {
		var pointValueInfo model.PointValue
		trans.CopyFields(pointDetails[i], &pointValueInfo)
		if err := tx.Table(common.PointValue.TableName()).Create(&pointValueInfo).Error; err != nil {
			log.Error("新增流程节点值信息出错")
			panic("服务器错误")
		}
		pointDetails[i].Id = pointValueInfo.Id
		// 新增节点参与者,结束节点除外
		if pointDetails[i].PointType != 6 {
			pointDetails[i].ParticipantInfos = participantRepo.BatchInsertParticipantInfo(tx, pointValueInfo.Id, pointDetails[i].ParticipantInfos)
		}

	}
	return pointDetails
}

// SelectPointValueListByCondition 根据条件查询节点值列表
func (m PointValueRepositoryImpl) SelectPointValueListByCondition(db *gorm.DB, conditionMap map[string]interface{}) (pointValueList []model.PointValue) {
	db = db.Table(common.PointValue.TableName())
	for k, v := range conditionMap {
		db = db.Where(k, v)
	}
	db.Select(common.PointValue.GetAllColumn()).Scan(&pointValueList)
	return
}

// SelectPointValueInfoByPointIdAndApprovalCode 根据节点id和审批编码查询指定的节点值信息
func (m PointValueRepositoryImpl) SelectPointValueInfoByPointIdAndApprovalCode(db *gorm.DB, pointId int64, approvalCode string) (pointValueInfo model.PointValue) {
	//TODO implement me
	panic("implement me")
}

func (m PointValueRepositoryImpl) UpdatePointValueInfo(tx *gorm.DB, details []api_model.PointDetail) bool {
	for _, detail := range details {
		var pointValueInfo model.PointValue
		trans.CopyFields(detail, &pointValueInfo)
		resultStatus := participantRepo.UpdateParticipantInfo(tx, detail.ParticipantInfos)
		if !resultStatus {
			return false
		}
		if err := tx.Table(common.PointValue.TableName()).Updates(pointValueInfo).Error; err != nil {
			log.Error("更新节点信息错误: " + err.Error())
			return false
		}
	}
	return true
}

// DeletePointValueInfo 删除节点值数据
func (m PointValueRepositoryImpl) DeletePointValueInfo(tx *gorm.DB, details []api_model.PointDetail) bool {
	for _, detail := range details {
		resultStatus := participantRepo.DeleteParticipantInfo(tx, detail.ParticipantInfos)
		if !resultStatus {
			return false
		}
		if err := tx.Table(common.PointValue.TableName()).Where("id = ?", detail.Id).Delete(&common.PointValue).Error; err != nil {
			log.Error("删除节点信息错误: " + err.Error())
			return false
		}
	}
	return true
}

func (m PointValueRepositoryImpl) SelectPointValueInfoByApprovalCode(db *gorm.DB, approvalCode string) (nodeList []api_model.PointDetail) {
	var pointDetailList []api_model.PointValueDetail
	db.Table(common.PointValue.TableName()+" pv").Joins("LEFT JOIN "+common.PanelPoint.TableName()+" p on p.id=pv.point_id").
		Where("approval_code = ?", approvalCode).Select(common.PointValue.GetAllColumWithAlias("pv") + ",p.point_type,p.examine_type").Scan(&pointDetailList)
	nodeList = make([]api_model.PointDetail, len(pointDetailList))
	trans.DeepCopy(pointDetailList, &nodeList)
	for i := 0; i < len(nodeList); i++ {
		nodeList[i].ParticipantInfos = participantRepo.SelectParticipantListByPointValueId(db, nodeList[i].Id, nil)
		pointInfo := pointRepo.SelectSinglePanelPointInfoById(lib.ObtainCustomDb(), nodeList[i].PointId)
		nodeList[i].ExamineType = pointInfo.ExamineType
		nodeList[i].PointType = pointInfo.PointType
		nodeList[i].PointName = pointInfo.PointName
	}
	return nodeList
}
