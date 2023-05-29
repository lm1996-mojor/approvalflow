package process

import (
	"five.com/lk_flow/model"
)

// QueryListMdl 条件分页列表请求条件封装结构体
type QueryListMdl struct {
	Search  string `json:"search,omitempty"`  // 综合搜索项（名称）
	GroupId int64  `json:"groupId,omitempty"` // 分组id
}

//响应

// ListProcessDetail 流程响应列表结构体
type ListProcessDetail struct {
	model.ProcessGroups                 //流程分组信息
	ProcessList         []model.Process `json:"processList"` //该分组下的流程列表
}

type DetailProcess struct {
	model.Process
	GroupName string `json:"groupName"` //分组名称
}
