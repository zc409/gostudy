package main

import (
	"fmt"
)

//匿名字段
//字段比较少也比较简单的场景
//不常用！！！
// type person struct {
// 	string
// 	int
// }

// func main() {
// 	p1 := person{
// 		"小明",
// 		18,
// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p1.string)
// 	fmt.Println(p1.int)
// }

//结构体嵌套
// type address struct {
// 	province string
// 	city     string
// }

// type person struct {
// 	name string
// 	age  int
// 	addr address
// }

// type company struct {
// 	name string
// 	addr address
// }

// func main() {
// 	p1 := person{
// 		name: "小明",
// 		age:  18,
// 		addr: address{
// 			province: "hubei",
// 			city:     "wuhan",
// 		},
// 	}
// 	fmt.Println(p1)
// 	fmt.Printf("name:%s,age:%d,province:%s,city:%s\n", p1.name, p1.age, p1.addr.province, p1.addr.city)
// }

//匿名嵌套结构体 及 匿名嵌套结构体字段冲突
type address struct {
	province string
	city     string
}

type workplace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
	//address: address
	workplace
}

func main() {
	s1 := person{
		name: "小红",
		age:  18,
		address: address{
			province: "湖北",
			city:     "武汉",
		},
		workplace: workplace{
			province: "浙江",
			city:     "杭州",
		},
	}
	fmt.Printf("姓名：%s,省份：%s\n", s1.name, s1.address.city)   //现在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找
	fmt.Printf("姓名：%s,省份：%s\n", s1.name, s1.workplace.city) //匿名嵌套结构体字段冲突，必须指定匿名字段名
}
