package main

import (
	"fmt"
)

func main() {
	var s1 = new(int)
	var s2 = make(map[string]int, 3)
	*s1 = 10
	s2["wuhan"] = 1
	s2["cq"] = 2
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}
