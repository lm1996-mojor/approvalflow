package engine

import (
	"sort"

	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/utils/trans"
)

// ApprovalPointChangeFun 审批节点处理方法
func ApprovalPointChangeFun(cachePointDetails []api_model.PointDetail, changePointDetails api_model.PointDetail) []api_model.PointDetail {
	for i := 0; i < len(cachePointDetails); i++ {
		if cachePointDetails[i].Id == changePointDetails.Id {

			// 发起人节点通过，发起人通过.发起人节点中只有一个参与人（必定条件）
			if changePointDetails.PointType == 5 {
				changePointDetails.ParticipantInfos[0].ApprovalResult = 1
				changePointDetails.ParticipantInfos[0].Opinions = "发起人节点自行通过"
			} else {
				/*
					审批形式（1 会签 2 或签）：ExamineType
						1  会签		所有参与人都开始审核，单必须是所有的参与人审批都通过才可以进行，节点进度修改
						2  或签		任意一个参与人审批通过，则修改节点的进度为通过
				*/
				// TODO: 发送站内消息 -- 消息系统暂未开发（多应用，多租户）

				// 设置当前节点为当前处理节点
				changePointDetails.IsCurrentPoint = true
				for j := 0; j < len(changePointDetails.ParticipantInfos); j++ {
					// 修改节点中参与人的状态
					changePointDetails.ParticipantInfos[j].ApprovalResult = 4
				}
			}
			// 修改数据源中的当前节点数据
			cachePointDetails[i] = changePointDetails
			break
		}
	}
	return cachePointDetails
}

// NextStepPoint 找到当前节点的下一个节点,并调整节点审批状态及其参与人审批状态,包含消息发送
func NextStepPoint(srcPointData []api_model.PointDetail, currentPointInfo api_model.PointDetail) []api_model.PointDetail {
	for i := 0; i < len(srcPointData); i++ {
		if srcPointData[i].Id == currentPointInfo.Id {
			// 将当前节点的信息同步至数据源中
			trans.DeepCopy(currentPointInfo, &srcPointData[i])
		}
		// 寻找当前节点的下一个节点
		if currentPointInfo.NextStep == srcPointData[i].PointId {
			// 重点先处理审批节点和抄送节点
			// PointRate	|节点审批进度 |（1 同意 2退回 3驳回 4审批中 5待执行 6撤销）
			// 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
			// 其他类型暂时未处理
			switch srcPointData[i].PointType {
			case 1:

				// 修改节点状态，并根据节点的审批类型来设置节点参与人的状态
				srcPointData[i].PointRate = 4
				srcPointData = ApprovalPointChangeFun(srcPointData, srcPointData[i])
			case 2:
				// 如果为抄送人节点，则先发送消息给抄送人,且该节点的进度为通过。再进行递归重新找节点直到找出审批人节点
				// TODO: 发送站内消息 -- 消息系统暂未开发
				srcPointData[i].PointRate = 1
				for j := 0; j < len(srcPointData[i].ParticipantInfos); j++ {
					// 设置抄送人审批结果为通过
					srcPointData[i].ParticipantInfos[j].ApprovalResult = 1
					srcPointData[i].ParticipantInfos[j].Opinions = "抄送人信息已送达"
				}
				srcPointData = NextStepPoint(srcPointData, srcPointData[i])
			case 6:
				// 设置结束节点的数据
				srcPointData[i].IsCurrentPoint = true
				srcPointData[i].PointRate = 1

			default:
				panic("找不到对应的节点类型")
			}
			// 设置srcPointData[i]上一个节点的是否为当前审批节点的状态为：否
			for k := 0; k < len(srcPointData); k++ {
				if srcPointData[i].PointId == srcPointData[k].NextStep {
					srcPointData[k].IsCurrentPoint = false
				}
			}
			break
		}
	}
	return srcPointData
}

// CheckPointIsApproval 检查指定节点是否已经审批
func CheckPointIsApproval(srcPointDataList []api_model.PointDetail, pointValueId int64) bool {
	for _, pointData := range srcPointDataList {
		if pointData.Id == pointValueId {
			for _, participantInfo := range pointData.ParticipantInfos {
				if participantInfo.ApprovalResult > 3 {
					return true
				}
			}
		}
	}
	return false
}

// ResetBatchPointApprovalStatus 重置指定批量节点的节点进度，并重置该节点中所有的参与人进度
func ResetBatchPointApprovalStatus(pointDetails []api_model.PointDetail, ids []int64) []api_model.PointDetail {
	for _, id := range ids {
		for i := 0; i < len(pointDetails); i++ {
			if pointDetails[i].Id == id {
				pointDetails[i].IsCurrentPoint = false
				pointDetails[i].PointRate = 5
				pointDetails[i].ParticipantInfos = ResetParticipantsApprovalResult(pointDetails[i].ParticipantInfos, 5)
			}
		}
	}
	return pointDetails
}

// ResetSinglePointApprovalStatus 重置单个指定节点的节点进度，并重置该节点中所有的参与人进度
func ResetSinglePointApprovalStatus(pointDetails []api_model.PointDetail, id int64) []api_model.PointDetail {
	for i := 0; i < len(pointDetails); i++ {
		if pointDetails[i].Id == id {
			pointDetails[i].IsCurrentPoint = true
			pointDetails[i].PointRate = 4
			pointDetails[i].ParticipantInfos = ResetParticipantsApprovalResult(pointDetails[i].ParticipantInfos, 4)
		}
	}
	return pointDetails
}

// FindThePreviousApprovalNodeOfTheSpecifiedNode 找出指定节点的上一个审批节点中间经过的所有id
func FindThePreviousApprovalNodeOfTheSpecifiedNode(appointedNodeId int64, srcPointData []api_model.PointDetail) (previousApprovalNodeId []int64) {
	//
	for _, datum := range srcPointData {
		for _, pointDetail := range srcPointData {
			if pointDetail.Id == appointedNodeId {
				if datum.NextStep == pointDetail.PointId {
					// 如果该节点不是审批节点则继续往上找
					if datum.PointType != 1 {
						previousApprovalNodeId = append(previousApprovalNodeId, FindThePreviousApprovalNodeOfTheSpecifiedNode(datum.Id, srcPointData)...)
					}
					previousApprovalNodeId = append(previousApprovalNodeId, datum.Id)
					break
				}
			}
		}

	}
	return previousApprovalNodeId
}

// FindPreviousApprovalId 找出指定节点的上一个审批节点Id
func FindPreviousApprovalId(appointedNodeId int64, srcPointData []api_model.PointDetail) (previousApprovalNodeId int64) {
	for _, datum := range srcPointData {
		for _, pointDetail := range srcPointData {
			if pointDetail.Id == appointedNodeId {
				if datum.NextStep == pointDetail.PointId {
					// 如果该节点不是审批节点则继续往上找
					if datum.PointType != 1 {
						previousApprovalNodeId = FindPreviousApprovalId(datum.Id, srcPointData)
					} else {
						return datum.Id
					}
				}
			}
		}
	}
	return previousApprovalNodeId
}

// ObtainCurrentPointValue 获取指定审批中的当前所处审批节点
func ObtainCurrentPointValue(srcPointData []api_model.PointDetail) (currentNodeInfo api_model.PointDetail) {
	for _, datum := range srcPointData {
		if datum.IsCurrentPoint {
			return datum
		}
	}
	log.Error("获取指定审批中的当前所处审批节点失败，没有找到当前节点")
	panic("服务器错误")
}

// ObtainPointValueInfo 获取指定节点信息
func ObtainPointValueInfo(srcPointData []api_model.PointDetail, pointValueId int64) (nodeInfo api_model.PointDetail) {
	for _, datum := range srcPointData {
		if datum.Id == pointValueId {
			nodeInfo = datum
		}
	}
	return nodeInfo
}

// DropAllNodesExceptTheInitiatorNode 删除流程中除发起人以外的节点
func DropAllNodesExceptTheInitiatorNode(srcPointData []api_model.PointDetail) (newPointDetailList []api_model.PointDetail) {
	for _, datum := range srcPointData {
		if datum.PointType != 5 {
			srcPointData = make([]api_model.PointDetail, 0)
			srcPointData = append(srcPointData, datum)
			return srcPointData
		}
	}
	log.Error("该流程不存在发起人节点")
	panic("该流程不存在发起人节点")
}

// CheckCurrentApprovalNodeIsContainUser 获取当前用户待处理的节点列表
func CheckCurrentApprovalNodeIsContainUser(pointValueList []api_model.PointDetail, userId int64) bool {
	for _, detail := range pointValueList {
		if detail.IsCurrentPoint {
			if CheckParticipantListContainUser(detail.ParticipantInfos, userId) {
				return true
			}
			break
		}
	}
	return false
}

// CheckIfThereAreAnyProcessedItemsForTheUserInTheApproval 检查该审批中是否存在该用户已处理的
func CheckIfThereAreAnyProcessedItemsForTheUserInTheApproval(pointValueList []api_model.PointDetail, userId int64) bool {
	for _, detail := range pointValueList {
		if detail.PointType == 1 {
			if DetermineWhetherTheUserHasBeenApproved(detail.ParticipantInfos, userId) {
				return true
			}
		}
	}
	return false
}

// CheckIfThereAreAnyReceivedByTheUserInTheApproval 获取当前用户收到的节点列表
func CheckIfThereAreAnyReceivedByTheUserInTheApproval(pointValueList []api_model.PointDetail, userId int64) bool {
	for _, detail := range pointValueList {
		if detail.PointType == 2 {
			if CheckParticipantListContainUser(detail.ParticipantInfos, userId) {
				if detail.PointRate == 1 {
					return true
				}
				// 已经找到但是没有发送过消息,没有必要再去找
				return false
			}
		}
	}
	return false
}

// 排序节点参与人
func sortParticipant(participants []model.Participant, sortStr string) (sortParticipantList []model.Participant) {
	sort.Slice(participants, func(i, j int) bool {
		if sortStr == "desc" {
			return participants[i].OrderNo > participants[j].OrderNo
		} else {
			return participants[i].OrderNo < participants[j].OrderNo
		}
	})
	return participants
}
