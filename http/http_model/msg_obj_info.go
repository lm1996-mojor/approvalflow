package http_model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

// MsgObjInfo 消息对象信息结构体
type MsgObjInfo struct {
	req.CommonModel
	ObjId          int64  `json:"objId,omitempty"`          // 对象id（系统/人员）
	ObjName        string `json:"objName,omitempty"`        // 对象名称
	ObjOrientation int8   `json:"objOrientation,omitempty"` // 对象定位（1 发送者 2 接受者）
	MsgId          int64  `json:"msgId,omitempty"`          // 消息id
	IsRead         int8   `json:"isRead,omitempty"`         // 是否已读（1 是 2 否）
}
