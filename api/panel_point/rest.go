package panel_point

import (
	"five.com/lk_flow/model"
	"five.com/lk_flow/model/rel"
)

// QueryListMdl 条件分页列表请求条件封装结构体
type QueryListMdl struct {
	ProcessId int64 `json:"processId"` // 流程id
}

//响应

type PanelPointDetail struct {
	model.PanelPoint
	RelMdl []rel.PointRelDesc `json:"relDescDetail"` //关系描述数组
}
