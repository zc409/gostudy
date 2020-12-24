package main

import "fmt"

func main() {
	//切片的定义
	var s1 []int
	var s2 []string
	fmt.Printf("%T,%T\n", s1, s2)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "武汉"}
	//长度和容量
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d,cap:%d\n", len(s2), cap(s2))
	//由数组得到切片
	s3 := [...]int{1, 2, 3, 4, 5}
	s4 := s3[0:3]
	fmt.Printf("s4 is:%v\n", s4)
	//容量是从切片第一个元素到数组最后一个元素的个数，长度是元素个数
	fmt.Printf("len is %d,cap is %d\n", len(s4), cap(s4))
	//切片在切片
	s5 := s4[1:]
	fmt.Printf("s4 is %v;s5 is %v\n", s4, s5)
	s3[1] = 1300
	fmt.Printf("s4 is %v;s5 is %v\n", s4, s5)
}
