package main

//构造函数
import "fmt"

type person struct {
	name   string
	gender string
	age    int
}

//构造函数：约定俗成用new开头
//返回的是结构体还是结构体指针是要根据实际情况考虑，字段较少可以返回值，
//字段较多返回指针，减少程序运行的内存开销
// func newPerson(name string, gender string, age int) person {        //返回值
// 	return person{
// 		name:   name,
// 		gender: gender,
// 		age:    age,
// 	}
// }
func newPerson(name string, gender string, age int) *person { //返回指针
	return &person{
		name:   name,
		gender: gender,
		age:    age,
	}
}
func main() {
	s1 := newPerson("xiaohong", "male", 8)
	fmt.Printf("%#v\n", s1)
}
