package message_api

import (
	"five.com/lk_flow/http/http_model"
)

// MsgInfo 消息主体和消息对象组合结构体
type MsgInfo struct {
	http_model.MessageInfo
	MsgObjList []http_model.MsgObjInfo `json:"msgObjList"`
}
