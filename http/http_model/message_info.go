package http_model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

// MessageInfo 消息主体信息结构体
type MessageInfo struct {
	req.CommonModel
	AppCode     string `json:"appCode,omitempty"`        // 应用编码
	ClientId    int64  `json:"clientId,omitempty"`       // 租户id
	MsgTitle    string `json:"msgTitle,omitempty"`       // 消息标题
	Content     string `json:"msgContent,omitempty"`     // 消息主体
	ContentType int    `json:"msgContentType,omitempty"` // 消息主体类型（普通文本、富文本、单操作链接、操作链接+文本、附件+文本）
	MsgLevel    int8   `json:"msgLevel,omitempty"`       // 消息优先级(1、高 2、中 3、低)
	OwnerType   int    `json:"ownerType,omitempty"`      // 消息所属类型（根据应用编码改变而改变,例如：应用编码是审批系统，则该类型至少有审批类型消息）
}
