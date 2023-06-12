package engine

import (
	"five.com/lk_flow/model"
)

// ParticipantApproval 流程参与人审批信息执行
func ParticipantApproval(cacheParticipantInfos []model.Participant, changeParticipant model.Participant) []model.Participant {
	for i := 0; i < len(cacheParticipantInfos); i++ {
		if cacheParticipantInfos[i].ObjId == changeParticipant.ObjId {
			cacheParticipantInfos[i].ApprovalResult = changeParticipant.ApprovalResult
			if changeParticipant.Opinions != "" {
				cacheParticipantInfos[i].Opinions = changeParticipant.Opinions
			}
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
func ResetParticipantsApprovalResult(participantInfoList []model.Participant, rate int8) []model.Participant {
	//审批结果:（1 同意 2退回 3驳回 4审批中 5待执行 6无操作 7 撤销）
	for i := 0; i < len(participantInfoList); i++ {
		participantInfoList[i].ApprovalResult = rate
		participantInfoList[i].Opinions = ""
	}
	return participantInfoList
}

// CheckWhetherTheParticipantHasBeenApproved 检查参与人是否已经审核
func CheckWhetherTheParticipantHasBeenApproved(participantInfoList []model.Participant) bool {
	//ApprovalResult 参与人审批状态 （1 同意 2退回 3驳回 4审批中 5待执行 6无操作 7 撤销）
	for _, participant := range participantInfoList {
		if participant.ApprovalResult != 5 {
			return true
		}
	}
	return false
}

// CheckParticipantListContainUser 检查参与列表是否包含指定用户
func CheckParticipantListContainUser(participantInfoList []model.Participant, userId int64) bool {
	for _, participant := range participantInfoList {
		if participant.ObjId == userId {
			return true
		}
	}
	return false
}

// DetermineWhetherTheUserHasBeenApproved 判断指定用户是否审批过
func DetermineWhetherTheUserHasBeenApproved(participantInfoList []model.Participant, userId int64) bool {
	for _, participant := range participantInfoList {
		if participant.ObjId == userId {
			if participant.ApprovalResult < 4 {
				return false
			}
		}
	}
	return true
}
