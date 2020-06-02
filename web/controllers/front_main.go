package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

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
//func (uc *FrontController) BeforeActivation(a mvc.BeforeActivation) {
//	a.Handle("GET", "/blog", "BlogGet", middleware.FrontMiddlewareHere)
//}

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

