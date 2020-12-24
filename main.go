package main

import "fmt"

func main() {
	s2 := [...]string{"武汉", "重庆", "成都"}
	for i, v := range s2 {
		fmt.Println(i, v)
		fmt.Printf("%T and %s\n", v, v)
	}
}
