package main

import "fmt"

//结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//狗类
type dog struct {
	feet   uint8
	animal //animal拥有的方法dog也拥有了（利用匿名结构体的语法糖）
}

//给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

func main() {
	s1 := dog{
		feet: 4,
		animal: animal{
			name: "小黑",
		},
	}
	s1.wang()
	s1.move()
	fmt.Printf("type:%T,value:%d", s1.feet, s1.feet)
}
