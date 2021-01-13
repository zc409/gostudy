package main

import "fmt"

type car interface {
	move()
}

type peugeot struct {
	name string
}

type ferari struct {
	name string
}

func (p peugeot) move() {
	fmt.Printf("%s run sportly\n", p.name)
}

func (f ferari) move() {
	fmt.Printf("%s run very fastly", f.name)
}

func main() {
	var c car      //声明一个car类型的变量
	p1 := peugeot{ //实例化一个标致
		name: "peugeot408",
	}
	f1 := ferari{ //实例化一个法拉利
		name: "ferari 458",
	}
	c = p1 //可以把标致实例直接赋值给c
	c.move()
	c = f1 //可以把法拉利实例直接赋值给c
	c.move()
}
