package main

import (
	"fmt"
	"time"
)

func main() {
	//程序开始时间
	st := time.Now()
	counts := 0
	//利用计时器每秒执行一次
	for i := range time.Tick(time.Second) {
		fmt.Println(i)
		counts++
		if counts > 3 {
			break
		}
	}
	//程序结束时间
	ed := time.Now()
	costtime := ed.Sub(st)
	//相差时间格式化为毫秒
	fmt.Println(costtime.Milliseconds())
}
