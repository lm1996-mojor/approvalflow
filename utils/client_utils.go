package utils

import (
	"fmt"

	_const "five.com/technical_center/core_library.git/const"
	"five.com/technical_center/core_library.git/store"
	"github.com/spf13/cast"
)

func GetClientId() int64 {
	//获取租户id
	var clientId int64
	value, ok := store.Get(fmt.Sprintf("%p", &store.PoInterKey) + _const.ClientID)
	if !ok {
		clientId = 0
	} else {
		clientId = cast.ToInt64(value)
	}
	return clientId
}
