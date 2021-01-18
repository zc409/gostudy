package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	minute := now.Minute()
	hour := now.Hour()
	second := now.Second()
	fmt.Printf("%v-%v-%0v %v:%v:%v\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println(now.Unix())             //获取时间戳
	fmt.Println(now.UnixNano())         //纳秒时间戳
	timetmp := time.Unix(1610523284, 0) //时间戳转化为时间类型，时间戳单位为秒
	fmt.Println(timetmp)
	yestd := now.Add(-24 * time.Hour) //时间转化为24小时前
	fmt.Println(yestd)
	dd := now.Sub(yestd) //now-yestd
	fmt.Println(dd)      //24h0m0s

	// tmpt := time.Tick(2 * time.Second) //两秒钟执行一次
	// for i := range tmpt {
	// 	fmt.Println(i)
	// }
	//24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05 Jan Mon"))
	//12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM"))

	//按照对应的格式解析字符串类型的时间
	timeobj, err := time.Parse("2006-01-02", "2021-01-13")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Println(timeobj)
	fmt.Println(timeobj.Unix())
}
