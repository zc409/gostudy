package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 //赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1) //copy
	fmt.Printf("a1=%v,a2=%v,a3=%v\n", a1, a2, a3)
	a1[0] = 100
	fmt.Printf("a1=%v,a2=%v,a3=%v\n", a1, a2, a3)
	//将切片中的元素删除
	x1 := [...]int{1, 2, 3}
	s1 := x1[:]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("x1=%v,len(x1)=%d,cap(x1)=%d\n", x1, len(x1), cap(x1))
}
