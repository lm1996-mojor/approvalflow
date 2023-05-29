package subset

import (
	"five.com/lk_flow/model"
)

// ListQueryMdl 分页列表条件查询结构体
type ListQueryMdl struct {
	ProcessId int64 `json:"processId"` // 流程id 必填
}

type SubDetail struct {
	model.Subset
}
