package main

import (
	"github.com/kataras/iris"
	"goBlog/utils"
	"goBlog/web/controllers"
)


func main()  {
	app := controllers.NewMvcApp()

	conf := utils.InitConfig()
	app.Run(
		iris.Addr(":" + conf.Port),
		iris.WithoutServerError(iris.ErrServerClosed),//无服务错误提示
		iris.WithOptimizations,
		iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
}



