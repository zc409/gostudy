package main

import "fmt"

type person struct {
	name   string
	gender string
	age    int
}

func main() {
	var test = new(person)
	fmt.Printf("%T\n", test)
	fmt.Printf("%v\n", test)
	fmt.Printf("%#v\n", test)
	fmt.Printf("%p\n", test)

	//key-value初始化
	var p1 = person{
		name:   "xiaoming",
		gender: "man",
	}
	fmt.Println(p1)

	//key-value列表形式初始化
	var p2 = person{
		"xiaohong",
		"male",
		9,
	}
	fmt.Println(p2)

	//结构体指针初始化
	var p3 = &person{
		name: "xiaobai",
	}
	fmt.Printf("%#v\n", p3)

	//结构体占用一块连续的内存空间
	type x struct {
		a int8
		b int8
		c int8
	}
	y := x{
		int8(1),
		int8(2),
		int8(3),
	}
	fmt.Printf("%p\n", &(y.a))
	fmt.Printf("%p\n", &(y.b))
	fmt.Printf("%p\n", &(y.c))
}
