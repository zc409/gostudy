package main

import "fmt"

//函数类型
func f1() {
	fmt.Println("helo world")
}

func f2() int {
	return 10
}

//函数也可以作为参数的类型
func t1() int {
	return 100
}

func t2(x int, y int) int {
	return x + y
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func t3(a, b int) int {
	return a + b
}

func f5(x func() int) func(int, int) int {
	return t3
}

func main() {
	a := f1
	b := f2
	fmt.Printf("%T   %T\n", a, b)
	f3(t1)
	//f3(t2)  te函数类型不符合无法传入
	c := f5(t1)          //c赋值返回的函数即t3
	fmt.Println(c(2, 3)) //5

}
