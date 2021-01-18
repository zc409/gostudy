package main

import (
	"time"

	"github.com/zc409/gostudy/day5/mylog"
)

func main() {
	for {
		log1 := mylog.Newloger("info")
		log1.Info("这是一个info级别的日志")
		log1.Warn("这是一个warn级别的日志")
		log1.Error("这是一个error级别的日志")
		log1.Fatal("这是一个fatal级别的日志")
		time.Sleep(2 * time.Second)
	}
}
