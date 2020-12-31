package main

import (
	"fmt"
)

//闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//f2无法直接放入f1中使用，定义一个函数对f2进行包装
func f3(f func(int, int), m, n int) func() {
	tem := func() {
		f(m, n)
	}
	return tem
}

func main() {
	tempf2 := f3(f2, 11, 22)
	f1(tempf2)
}
