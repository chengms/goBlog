package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)


func NewMvcApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")	//设置日志级别
	//可以从任何与 http 相关的panics（http-relative panics） 中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	app.Logger().Info("--start web--")

	tmp := iris.HTML("./views", ".html")
	//自动重载模板文件
	tmp.Reload(true)
	//加载静态文件
	app.StaticWeb("/static", "./static")

	//错误处理
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, InternalServerError)

	//注册进去
	app.RegisterView(tmp)

	//路由组的mvc处理
	mvc.Configure(app, Basic)

	return app
}

// 在路由组下面创建 context.Handle(new(Branch))
func Basic(context *mvc.Application) {
	iris.New().Logger().Info(" root 控制器")
	Root(context)

	iris.New().Logger().Info(" 定义 Front 子路径")
	Front(context)

	iris.New().Logger().Info(" 定义 Admin 子路径")
	Admin(context)
}


