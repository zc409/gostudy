package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的练习题
	var a = make([]int, 5, 10) //创建切片
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //  [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	//切片排序
	var a1 = [...]int{10, 1, 5, 3}
	sort.Ints(a1[:])
	fmt.Println(a1) //[1 3 5 10]
}
