package controllers

import (
	"github.com/kataras/iris"
)

//404
func NotFound(ctx iris.Context) {
	ctx.View("404.html")
}


//500
func InternalServerError(ctx iris.Context) {
	ctx.WriteString("凉凉,服务器错误")
}

