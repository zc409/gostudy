package main

import "fmt"

//defer多用于函数
func testdefer() {
	fmt.Println("start")
	defer fmt.Println("middle1") //defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("middle2") //一个函数中可以有多个defer语句
	defer fmt.Println("middle3") //多个defer语句按照后进先出的顺序延迟执行
	fmt.Println("end")
}

func main() {
	testdefer()
}
