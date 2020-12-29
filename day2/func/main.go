package main

import "fmt"

//函数的定义
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值的
func f3() int {
	return 3
}

//命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return后省略
}

//返回多个值
func f5() (int, int) {
	return 1, 2
}

//参数的类型简写
//当参数中连续多个参数的类型一致时
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x int, y ...int) { //y是切片类型
	fmt.Println(x)
	fmt.Println(y)
}

//go语言中函数没有默认参数这个概念
func main() {
	s1 := sum(10, 20)
	fmt.Println(s1)
	f1(1, 2)
	f2()
	fmt.Println(f3())
	x, y := f5()
	fmt.Printf("x=%d,y=%d\n", x, y)
	f7(1, 2, 3, 4, 5) //1  [2 3 4 5]
}
