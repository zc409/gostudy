package main

import (
	"fmt"
	"github.com/zc409/gostudy/day5/testpackage"
)

func main() {
	testpackage.Helo() //应用的package同函数首行定义的package xxx 相同
	//testpackage.Hehe()
	fmt.Println("main package")
}
