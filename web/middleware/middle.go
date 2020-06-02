package middleware

import "github.com/kataras/iris"


func RootMiddlewareHere(ctx iris.Context) {
	ctx.Application().Logger().Warnf("Root 调用中间件")
	ctx.Next()
}

func FrontMiddlewareHere(ctx iris.Context) {
	ctx.Application().Logger().Warnf("Front 调用中间件")
	ctx.Next()
}

func AdminMiddlewareHere(ctx iris.Context) {
	ctx.Application().Logger().Warnf("Admin 调用中间件")
	ctx.Next()
}
