package main

import "fmt"

func main() {
	var n = 100
	//打印类型
	fmt.Printf("%T\n", n)
	//打印值
	fmt.Printf("%v\n", n)
	//打印转换为二进制
	fmt.Printf("%b\n", n)
	//打印数字
	fmt.Printf("%d\n", n)
	//打印转换为八进制
	fmt.Printf("%o\n", n)
	//打印转换为十六进制
	fmt.Printf("%x\n", n)
	var s = "helo world!"
	//打印字符串
	fmt.Printf("%s\n", s)
	//打印值
	fmt.Printf("%v\n", s)
	//打印如果是字符加双引号
	fmt.Printf("字符串：%#v\n", s)
	fmt.Printf("数字：%#v\n", n)
}
