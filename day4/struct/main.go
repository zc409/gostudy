package main

import "fmt"

// type dog struct {
// 	gender string
// 	name   string
// 	age    int
// 	colour string
// 	hobby  []string
// }

// func main() {
// 	var hei dog
// 	hei.gender = "man"
// 	hei.name = "xiaohei"
// 	hei.age = 3
// 	hei.colour = "black"
// 	hei.hobby = []string{"eat", "sleep", "playball"}
// 	fmt.Printf("type:%T,value:%v", hei, hei)
// }

//匿名结构体:多用于临时场景
func main() {
	var test struct {
		name   string
		age    int
		gender string
		hobby  []string
	}
	test.name = "xiaoming"
	test.age = 12
	test.gender = "man"
	test.hobby = []string{"basketball", "football", "game"}
	fmt.Printf("type:%T,value:%v\n", test, test)
}
