package main

import (
	"five.com/technical_center/core_library.git/config"
	"five.com/technical_center/core_library.git/global"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	global.Init(app, 100)
	err := app.Run(iris.Addr(":"+config.Sysconfig.App.Port), iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutBodyConsumptionOnUnmarshal, iris.WithoutPathCorrectionRedirection)
	if err != nil {
		panic(err)
	}
}

//在terminal中运行一下命令
//git config --global url."http://gitlab.five.com:8888/technical_center/core_library.git".insteadOf https://five.com/technical_center/core_library
//go env -w GO111MODULE=on
//go env -w GOPROXY=https://goproxy.cn,direct
//go env -w GOPRIVATE=five.com/technical_center/core_library.git
//go get five.com/core_library.git
