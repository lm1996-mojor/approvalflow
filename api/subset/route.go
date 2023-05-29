package subset

import (
	"five.com/technical_center/core_library.git/global"
	"five.com/technical_center/core_library.git/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Init(app *iris.Application) {
	mvc.New(app.Party("/flow/subset")).Handle(NewController(NewMysqlRepository()))
	log.Info("初始化平台系统流程表单控件子集管理模块")
}

const runLevel = 10

func init() {
	global.RegisterInit(global.Initiator{Action: Init, Level: runLevel})
}
