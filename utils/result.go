package utils

import (
	"five.com/technical_center/core_library.git/rest"
)

// Result 统一返回-节省代码量
func Result(element interface{}, resultJsonName string) rest.Result {
	resultMap := make(map[string]interface{}, 1)
	resultMap[resultJsonName] = element
	return rest.SuccessResult(resultMap)
}
