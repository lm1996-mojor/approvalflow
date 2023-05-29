package condition_detail_groups

import (
	"five.com/lk_flow/model"
)

type QueryListMdl struct {
	ConditionInfoId int64 `json:"conditionInfoId,omitempty"` // 条件主体信息id
}

type GroupsDetail struct {
	model.ConditionDetailGroups
	ConditionDetailInfos []model.ConditionDetailInfo `json:"conditionDetailInfos,omitempty"`
}
