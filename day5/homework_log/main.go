package main

import (
	"time"

	"github.com/zc409/gostudy/day5/mylog"
)

func main() {
	var log1 mylog.Minelog
	for {
		//log1 = mylog.Newloger("info")
		//log1 = mylog.Newfileloger("info", "D:/gogo/src/github.com/zc409/gostudy/day5/homework_log/", "log.txt", 3*1024)
		log1 = mylog.Melog("info", "f")
		log1.Info("这是一个info级别的日志")
		log1.Warn("这是一个warn级别的日志")
		log1.Error("这是一个error级别的日志")
		log1.Fatal("这是一个fatal级别的日志")
		time.Sleep(2 * time.Second)
	}
}
