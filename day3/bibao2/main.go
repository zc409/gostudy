package main

import (
	"fmt"
)

//闭包是什么？
//闭包是一个函数，这个函数包含了他外部作用域的一个变量

//底层的原理：
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，现在自己内部找，找不到往外层找
// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	ret := adder(100)
// 	ret2 := ret(200)
// 	fmt.Println(ret2)
// }
//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。
//在f的生命周期内，变量x也一直有效。
// func adder2(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	var f = adder2(10)
// 	fmt.Println(f(10)) //20
// 	fmt.Println(f(20)) //40
// 	fmt.Println(f(30)) //70

// 	f1 := adder2(20)
// 	fmt.Println(f1(40)) //60
// 	fmt.Println(f1(50)) //110
// }
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	f1, f2 := calc(10)

	fmt.Println(f1(1), f2(2)) //11,9
	fmt.Println(f1(3), f2(4)) //12,8
	fmt.Println(f1(5), f2(6)) //13,7
}
