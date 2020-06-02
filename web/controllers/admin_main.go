package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goBlog/web/middleware"
)

type AdminController struct {
	// 上下文对象
	Ctx iris.Context

	// Session 对象
	Session *sessions.Session
}

func Admin(context *mvc.Application)  {
	context.Party("/admin").Handle(new(AdminController))
}

//中间件
func (uc *AdminController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/", "GetAdmin", middleware.AdminMiddlewareHere)
}

func (uc *AdminController)GetAdmin(ctx iris.Context) {

	ctx.SetCookieKV("Sub", "Sub") // <-- 设置一个Cookie

	if err := ctx.View("admin/admin.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

func (uc *AdminController)Post(ctx iris.Context) {

	ctx.Redirect("/admin")
}



