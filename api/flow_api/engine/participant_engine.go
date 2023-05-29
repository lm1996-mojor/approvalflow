package engine

import (
	"five.com/lk_flow/model"
)

// ParticipantApproval 流程参与人审批信息执行
func ParticipantApproval(cacheParticipantInfos []model.Participant, changeParticipant model.Participant) []model.Participant {
	for i := 0; i < len(cacheParticipantInfos); i++ {
		if cacheParticipantInfos[i].ObjId == changeParticipant.ObjId {
			cacheParticipantInfos[i].ApprovalResult = changeParticipant.ApprovalResult
		}
	}
	return cacheParticipantInfos
}

// CheckParticipantsApprovalResult 根据条件检查指定节点中的参与人的审批结果是否和指定的相同
func CheckParticipantsApprovalResult(participantInfoList []model.Participant, approvalResult int8, isAll bool) bool {
	for _, participant := range participantInfoList {
		if isAll {
			if participant.ApprovalResult != approvalResult {
				return false
			}
		} else {
			if participant.ApprovalResult == approvalResult {
				return true
			}
		}
	}
	if isAll {
		return true
	} else {
		return false
	}
}

// ResetParticipantsApprovalResult 重置的参与人状态和意见
func ResetParticipantsApprovalResult(participantInfoList []model.Participant) []model.Participant {
	for i := 0; i < len(participantInfoList); i++ {
		participantInfoList[i].ApprovalResult = 5
		participantInfoList[i].Opinions = ""
	}
	return participantInfoList
}

// CheckWhetherTheParticipantHasBeenApproved 检查参与人是否已经审核
func CheckWhetherTheParticipantHasBeenApproved(participantInfoList []model.Participant) bool {
	//ApprovalResult 参与人审批状态 （1 同意 2退回 3驳回 4审批中 5待执行 6无操作）
	for _, participant := range participantInfoList {
		if participant.ApprovalResult != 5 {
			return true
		}
	}
	return false
}
