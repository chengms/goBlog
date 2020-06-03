package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)


func NewMvcApp() *iris.Application {
	//Logf := logs.NewLogFile()
	//defer Logf.Close()

	app := iris.New()

	//app.Logger().SetOutput(io.MultiWriter(Logf, os.Stdout)).SetLevel("debug")	//设置日志级别
	app.Logger().SetLevel("debug")

	// recover 中间件从任何异常中恢复，如果有异常，则写入500状态码（服务器内部错误）
	app.Use(recover.New())
	//并将请求记录到终端。
	requestLogger := logger.New(logger.Config{
		// 是否开启状态码
		Status: true,
		// 是否记录远程IP地址
		IP: true,
		// 是否呈现HTTP谓词
		Method: true,
		// 是否记录请求路径
		Path: true,
		// 是否开启查询追加
		Query: true,

		// 若为空则从 `ctx.Values().Get("logger_message") 获取内容
		// 在其中添加日志
		MessageContextKeys: []string{"logger_message"},

		// 为空则从 `ctx.GetHeader("User-Agent") 获取头信息
		//MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	app.Logger().Info("--start web--")

	tmp := iris.HTML("./views", ".html")
	//自动重载模板文件
	tmp.Reload(true)
	//加载静态文件
	app.StaticWeb("/static", "./static")

	// 网站图标
	// GET: http://localhost:8080/favicon.ico
	app.Favicon("./static/img/favicon.ico")
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


