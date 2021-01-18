package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//获取上海的时区类型指针
	timeloc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("get location err,err:%v\n", err)
		return
	}
	//将utc字符串转化为上海时区的time类型变量
	yestd, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-15 08:37:00", timeloc)
	if err != nil {
		fmt.Printf("get yestd err,err:%v\n", err)
		return
	}
	res := yestd.Sub(now)
	//格式化为小时输出，保留2位小数点
	fmt.Printf("%.2f", res.Hours())
}
