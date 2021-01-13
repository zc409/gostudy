package main

import "fmt"

type animal interface {
	move()
	speak()
}

type dog struct {
	name  string
	color string
}

type cat struct {
	name  string
	color string
}

func (d dog) move() {
	fmt.Printf("%s color is %s,run adn run\n", d.name, d.color)
}

func (d dog) speak() {
	fmt.Printf("%s 汪汪汪\n", d.name)
}

func (c cat) move() {
	fmt.Printf("%s color is %s,run adn run\n", c.name, c.color)
}

func (c cat) speak() {
	fmt.Printf("%s 喵喵喵\n", c.name)
}

func mover(a animal) {
	a.move()
}

func speaker(a animal) {
	a.speak()
}

func main() {
	d1 := dog{
		name:  "小黑狗",
		color: "黑色",
	}
	c1 := cat{
		name:  "小花猫",
		color: "白黄",
	}
	mover(d1)
	mover(c1)
	speaker(d1)
	speaker(c1)
}
