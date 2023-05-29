package flow_api

import (
	"five.com/lk_flow/api/flow_api/controller"
	"five.com/lk_flow/api/flow_api/repository/repoImpl"
	"five.com/technical_center/core_library.git/global"
	"five.com/technical_center/core_library.git/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Init(app *iris.Application) {
	mvcApp := mvc.New(app.Party("/flow/open/api"))
	mvcApp.Handle(controller.ApprovalNewController(repoImpl.ProcessValueNewMysqlRepository()))
	log.Info("初始化审批开放Api模块")
}

const runLevel = 10

func init() {
	global.RegisterInit(global.Initiator{Action: Init, Level: runLevel})
}
