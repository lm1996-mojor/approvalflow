package engine

import (
	"sort"

	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/utils/trans"
)

// ApprovalPointChangeFun 启动流程节点，并调整节点中参与人的状态
func ApprovalPointChangeFun(cachePointDetails []api_model.PointDetail, changePointDetails api_model.PointDetail, approvalParam api_model.ApprovalParams) []api_model.PointDetail {
	for i := 0; i < len(cachePointDetails); i++ {
		if cachePointDetails[i].Id == changePointDetails.Id {
			// 发起人节点通过，发起人通过.发起人节点中只有一个参与人（必定条件）
			if changePointDetails.PointType == 5 {
				cachePointDetails[i].ParticipantInfos[0].ApprovalResult = 1
				cachePointDetails[i].ParticipantInfos[0].Opinions = "发起人节点自行通过"
			} else {
				/*
					审批形式（1 会签 2 或签）：ExamineType
						1  会签		所有参与人都开始审核，单必须是所有的参与人审批都通过才可以进行，节点进度修改
						2  或签		任意一个参与人审批通过，则修改节点的进度为通过
				*/
				// TODO: 发送站内消息 -- 消息系统暂未开发（多应用，多租户）

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
func NextStepPoint(srcPointData []api_model.PointDetail, currentPointInfo api_model.PointDetail, approvalParam api_model.ApprovalParams) []api_model.PointDetail {
	for i := 0; i < len(srcPointData); i++ {
		if srcPointData[i].Id == currentPointInfo.Id {
			// 将当前节点的信息同步至数据源中
			trans.DeepCopy(currentPointInfo, &srcPointData[i])
		}
		// 寻找当前节点的下一个节点
		if currentPointInfo.NextStep == srcPointData[i].PointId {
			// 重点先处理审批节点和抄送节点
			// PointRate	|节点审批进度 |（1 同意 2退回 3驳回 4审批中 5待执行 6无操作）
			// 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
			// 其他类型暂时未处理
			switch srcPointData[i].PointType {
			case 1:
				// 修改节点状态，并根据节点的审批类型来设置节点参与人的状态
				srcPointData[i].PointRate = 4
				srcPointData = ApprovalPointChangeFun(srcPointData, srcPointData[i], approvalParam)
			case 2:
				// 如果为抄送人节点，则先发送消息给抄送人。再进行递归重新找节点直到找出审批人节点
				// TODO: 发送站内消息 -- 消息系统暂未开发
				srcPointData[i].PointRate = 4
				srcPointData = NextStepPoint(srcPointData, srcPointData[i], approvalParam)
			default:
				panic("找不到对应的节点类型")
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

// ResetPointApprovalStatus 重置指定节点的节点进度，并重置该节点中所有的参与人进度
func ResetPointApprovalStatus(pointDetails []api_model.PointDetail, ids []int64) []api_model.PointDetail {
	//审批结果: 1 同意 2退回 3驳回 4审批中 5待执行 6无操作
	for _, id := range ids {
		for i := 0; i < len(pointDetails); i++ {
			if pointDetails[i].Id == id {
				pointDetails[i].PointRate = 5
				pointDetails[i].ParticipantInfos = ResetParticipantsApprovalResult(pointDetails[i].ParticipantInfos)
			}
		}
	}
	return pointDetails
}

// FindNeedChangeNodeIds 找出需要修改的节点ids并排除开始节点
func FindNeedChangeNodeIds(startPointId, endPointId, nextPointId int64, srcPointData []api_model.PointDetail) (ids []int64) {
	for _, datum := range srcPointData {
		if datum.Id == startPointId {
			if datum.NextStep == endPointId {
				ids = append(ids, datum.NextStep)
				return ids
			}
			var nextPoint api_model.PointDetail
			for _, pointDatum := range srcPointData {
				if pointDatum.Id == nextPointId {
					nextPoint = pointDatum
				}
			}
			ids = append(ids, nextPoint.Id)
			ids = append(ids, FindNeedChangeNodeIds(nextPoint.Id, endPointId, nextPoint.NextStep, srcPointData)...)
		}
	}
	return ids
}

// FindThePreviousApprovalNodeOfTheSpecifiedNode 找出指定节点的上一个审批节点
func FindThePreviousApprovalNodeOfTheSpecifiedNode(appointedNodeId int64, srcPointData []api_model.PointDetail) (previousApprovalNodeId []int64) {
	for _, datum := range srcPointData {
		if datum.NextStep == appointedNodeId {
			if datum.PointType != 1 {
				previousApprovalNodeId = append(previousApprovalNodeId, FindThePreviousApprovalNodeOfTheSpecifiedNode(datum.Id, srcPointData)...)
			}
			previousApprovalNodeId = append(previousApprovalNodeId, datum.Id)
			break
		}
	}
	return previousApprovalNodeId
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
