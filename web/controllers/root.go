package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goBlog/web/middleware"
	"time"
)

type RootController struct {
	// 上下文对象
	Ctx iris.Context

	// Session 对象
	Session *sessions.Session
}

func Root(context *mvc.Application)  {
	// 相当于mvc.New(app).Handle(new(UserController))
	context.Handle(new(RootController))
}

//中间件
func (uc *RootController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/", "IndexGet", middleware.RootMiddlewareHere)
}


func (uc *RootController) IndexGet() {
	iris.New().Logger().Info(" Get请求-->首页 ")

	// 绑定： {{ contents }}　为　"Hello world!"
	uc.Ctx.ViewData("datas", "Hello world!")

	now := time.Now().Format(uc.Ctx.Application().ConfigurationReadOnly().GetTimeFormat())
	uc.Ctx.ViewData("now", now)

	if err := uc.Ctx.View("index.html"); err != nil {
		uc.Ctx.Application().Logger().Infof(err.Error())
	}
}

func (uc *RootController) GetAbout() {
	iris.New().Logger().Info(" get请求-->about")

	if err := uc.Ctx.View("about.html"); err != nil {
		uc.Ctx.Application().Logger().Infof(err.Error())
	}
}

func (uc *RootController)GetLogin(ctx iris.Context)  {

	if err := ctx.View("admin/login.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

