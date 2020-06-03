package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"goBlog/web/middleware"
)

type FrontController struct {
	// 上下文对象
	Ctx iris.Context

	// Session 对象
	//Session *sessions.Session
}

func Front(context *mvc.Application)  {
	context.Party("/article").Handle(new(FrontController))
}


//中间件
func (uc *FrontController) BeforeActivation(front mvc.BeforeActivation) {
	// 设置子路由
	rut := front.Router()
	{
		rut.Get("/FrontTest", FrontTest)

	}

	front.Handle("GET", "/", "Get", middleware.FrontMiddlewareHere)
}

func (uc *FrontController)Get(ctx iris.Context) {

	if err := ctx.View("articles.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

func (uc *FrontController)Post(ctx iris.Context) {

	if err := ctx.View("articles.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

func FrontTest(ctx iris.Context) {
	iris.New().Logger().Info("Front Test")

	ctx.HTML(`"Front Test"`)
}

