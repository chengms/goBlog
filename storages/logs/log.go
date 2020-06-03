package logs

import (
	"os"
	"time"
)

// 以日期为文件名, 此处日期是个语法糖，只用了该日期的语法格式，并非指2006年.
func todayFilename() string {
	today := time.Now().Format("Jan-02-2006")
	return "./logs/" + today + ".txt"
}

func NewLogFile() *os.File {
	filename := todayFilename()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
