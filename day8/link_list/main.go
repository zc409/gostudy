package main

//如何判断一个链表有没有闭环

import (
	"fmt"
)

func setp() int {
	x := 1
	y := 2
	var n int
	fmt.Print("please input step number: ")
	//	fmt.Scanf("%d",&n)
	fmt.Scanln(&n)
	if n == 1 {
		return x
	} else if n == 2 {
		return y
	} else {
		for i := 0; i < n-2; i++ {
			x, y = y, x+y
		}
		return y
	}
}

func main() {
	num := setp()
	fmt.Println(num)
}
