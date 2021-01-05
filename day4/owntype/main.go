package main

import "fmt"

//自定义类型和类型别名

//type后面跟的是类型
type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) //自定义类型编译后还是自定义名  main.myInt

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m) //类型别名编译后还原成原名 int

	var c rune //rune是系统自带的类型别名
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c) //int32
}
