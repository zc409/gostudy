package main

import "fmt"

func main() {
	ss := [...]string{"北京", "上海", "广州"}
	for _, v := range ss {
		fmt.Println(v)
	}
	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}
}
