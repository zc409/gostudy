package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string, age int) person {
// 	return person{
// 		name: name,
// 		age:  age,
// 	}
// }

// func (p person) older() {
// 	p.age++
// }
// func (p *person) oldert() {
// 	//(*p).age++
// 	p.age++ //语法糖
// }
// func main() {
// 	s1 := newPerson("xiaoming", 18)
// 	s1.older()
// 	fmt.Println(s1)
// 	//(&s1).oldert()
// 	s1.oldert()
// 	fmt.Println(s1)
// }

//不能给别的包里面的类型添加方法，只能给自己包里面的类型添加方法
type myInt int

func (m *myInt) biger() {
	*m++
}

func main() {
	s1 := myInt(10)
	fmt.Printf("type:%T,value:%v\n", s1, s1)
	//(&s1).biger()
	s1.biger() //语法糖
	fmt.Printf("type:%T,value:%v\n", s1, s1)
}
