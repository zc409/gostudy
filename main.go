package main

import (
	"fmt"
)

func main() {
	hui := true
	s1 := "你好heoeh好你"
	s2 := []rune(s1)
	fmt.Printf("s2=%v\n", s2)
	s3 := len(s2) / 2
	for i := 0; i < s3; i++ {
		if s2[i] != s2[len(s2)-1-i] {
			hui = false
		}
	}
	if hui == true {
		fmt.Println("this is hui")
	} else {
		fmt.Println("this is not hui")
	}
}
