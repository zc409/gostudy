package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

//go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "male" //修改的是副本的gender
}

func f2(x *person) { //传入内存地址可以强行修改
	(*x).gender = "male" //根据内存地址找到那个原变量
	//x.gender = "male"  语法糖(*x)可以简写成x
}

func main() {
	var p person
	p.name = "xiaoming"
	p.gender = "man"
	f(p)
	fmt.Println(p.gender) //man
	f2(&p)
	fmt.Println(p.gender) //male
}
