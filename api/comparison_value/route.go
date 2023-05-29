package comparison_value

import (
	"five.com/technical_center/core_library.git/global"
	"five.com/technical_center/core_library.git/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Init(app *iris.Application) {
	mvc.New(app.Party("/flow")).Handle(NewController(NewMysqlRepository()))
	log.Info("初始化审批系统比较值信息管理管理模块")
}

const runLevel = 10

func init() {
	global.RegisterInit(global.Initiator{Action: Init, Level: runLevel})
}
