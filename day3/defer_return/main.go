package main

import (
	"fmt"
)

//go语言中函数的return不是原子操作，在底层是分为两部来执行
//第一步：返回值赋值
//第二部：真正的RET返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	x = 2
	defer func(x int) {
		x++
		fmt.Printf("funcx=%d\n", x)
	}(x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4()) //输出funcx=3,5。说明defer延迟只延迟里面的函数，不延迟函数赋值。
}
