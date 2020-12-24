package main

import "fmt"

func main() {
	s2 := "你好"
	s1 := []rune(s2)
	fmt.Printf("%T,%c\n", s1, s1)
	s3 := [...]string{"你", "好"}
	fmt.Printf("%T,%v", s3, s3)

	
}
