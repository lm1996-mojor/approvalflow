package examine_rule

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type QueryPageMdl struct {
	req.PageParam
	Search string `json:"search"` //综合搜索项（名称）
}
