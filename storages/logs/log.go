package logs

import "github.com/kataras/iris"

func Log(msg interface{})  {
	iris.New().Logger().Info(msg)
}
