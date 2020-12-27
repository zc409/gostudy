package main

import "fmt"

func main() {
	// s := [...]int{1, 2, 3}
	// fmt.Printf("%v\n", &s[0])
	// d := &s[0]
	// fmt.Printf("%v,%T\n", d, d) //指针类型是*int或者*string等
	// e := *d
	// fmt.Printf("%v,%T\n", e, e)
	// var s1 int
	// fmt.Printf("s1'point=%v\n", &s1)
	// fmt.Printf("s1=%v\n", *&s1)
	// s1 = 2
	// fmt.Printf("s1'point=%v\n", &s1)
	// fmt.Printf("s1=%v\n", *&s1)
	// make 和 new
	// var a1 *int //nil pointer
	// fmt.Println(a1)
	// var a2 = new(int) //new函数申请一个内存地址
	// fmt.Println(a2)   //内存地址
	// fmt.Println(*a2)  //0
	// *a2 = 100
	// fmt.Println(*a2) //100
	var b map[string]int
	b = make(map[string]int, 10)
	fmt.Printf("%v\n", b)
	b["测试一下"] = 100
	fmt.Println(b)
	fmt.Printf("len(b)=%d", len(b))
}
