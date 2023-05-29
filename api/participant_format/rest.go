package participant_format

import (
	"five.com/technical_center/core_library.git/rest/req"
)

// QueryListMdl 条件分页列表请求条件封装结构体
type QueryListMdl struct {
	req.PageParam
	Search string `json:"search"` // 综合搜索项（名称）
}

//响应
