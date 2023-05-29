package control_info

import (
	"five.com/lk_flow/model"
)

// ListQueryMdl 分页列表条件查询结构体
type ListQueryMdl struct {
	ParentId  int64 `json:"parentId,omitempty"`  //父级id（关联子集表和规则表id）必填
	OwnerType uint8 `json:"ownerType,omitempty"` // 控件所属类型（1 流程主体 2 流程规则）
}

type CtlDetail struct {
	model.ControlInfo
}
