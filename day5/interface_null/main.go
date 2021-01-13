package main

import "fmt"

//空接口
//interface：关键字
//interface{}：空接口类型

//空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "mike"
	m1["age"] = 18
	m1["merried"] = false
	m1["hobby"] = [...]string{"sing", "dance", "rap"}
	fmt.Println(m1) //空接口可以存储任何类型的变量

	show(false)
	show(nil)
	show(m1)
}
