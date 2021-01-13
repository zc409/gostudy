package main

import "fmt"

type student interface {
	working()
}

type littles struct {
	name string
}

type univers struct {
	name string
}

func (l *littles) working() { //方法使用指针接收者，实例化后传给接口变量只能传指针
	fmt.Printf("%s working very easy\n", l.name)
}

func (u univers) working() { //方法使用值接收者实例化后传给接口变量可以传值也可以传指针
	fmt.Printf("%s working widely\n", u.name)
}

func main() {
	var s student
	l1 := littles{
		name: "小学生",
	}
	u1 := &univers{ //实例化指针
		name: "大学生",
	}
	//s = l1        littles的方法是指针接收者，不能传值
	s = &l1 //littles的方法是指针接收者，只能传指针
	s.working()
	s = u1 //传指针可以
	fmt.Printf("%T\n", s)
	(*u1).working()
	fmt.Printf("%T\n", *u1)
	s.working()
	s = *u1 //传值也可以
	fmt.Printf("%T\n", s)
	s.working()
}
