package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// name := flag.String("nm", "小明", "请输入名字")
	// age := flag.Int("age", 18, "请输入年龄")
	// married := flag.Bool("ma", false, "结婚了吗")
	// ctime := flag.Duration("ct", time.Hour, "结婚多久了")
	var name string
	var age int
	var married bool
	var ctime time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&ctime, "d", 0, "时间间隔")
	flag.Parse()
	// fmt.Printf("type:%T,value:%#v\n", *name, *name)
	// fmt.Printf("type:%T,value:%#v\n", *age, *age)
	// fmt.Printf("type:%T,value:%#v\n", *married, *married)
	// fmt.Printf("type:%T,value:%#v\n", *ctime, *ctime)
	// fmt.Println(*ctime)
	fmt.Printf("type:%T,value:%#v\n", name, name)
	fmt.Printf("type:%T,value:%#v\n", age, age)
	fmt.Printf("type:%T,value:%#v\n", married, married)
	fmt.Printf("type:%T,value:%#v\n", ctime, ctime)
	fmt.Println(ctime)
	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}

// D:\gogo\src\github.com\zc409\gostudy\day8\flag_demo>go run main.go -name=test -age=19 a b c
// type:string,value:"test"
// type:int,value:19
// type:bool,value:false
// type:time.Duration,value:0
// 0s
// [a b c]
// 3
// 2