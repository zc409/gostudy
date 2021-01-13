package main

import "fmt"

//同一个结构体可以实现多个接口
//接口还可以嵌套

type animal interface { //接口嵌套
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步。。。")
}

// cat 实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func eating(e eater, food string) { //接口eater的函数
	e.eat(food)
}

func animaleat(a animal, food string) { //接口animal的函数
	a.eat(food)
}

func main() {
	c1 := &cat{
		name: "miaomiao",
		feet: 4,
	}
	eating(c1, "fish")      //cat的实例可以实现接口eater的函数
	animaleat(c1, "fisher") //cat的实例可以实现接口animal的函数
}
