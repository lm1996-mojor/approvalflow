package process_group

import (
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest/req"
)

type QueryPageMdl struct {
	req.PageParam
	Search   string `json:"search,omitempty"`  //综合搜索项（名称）
	AppCode  string `json:"appCode,omitempty"` // 应用编码
	ClientId int64  `json:"-,omitempty"`       // 租户id
}
type QueryListMdl struct {
	AppCode  string `json:"appCode,omitempty"` // 应用编码
	ClientId int64  `json:"-,omitempty"`       // 租户id
}

type ProcessGroupsDetail struct {
	model.ProcessGroups
	ProcessList []model.Process `json:"processList"` //该分组下的所有流程
}
