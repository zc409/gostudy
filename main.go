package main

import (
	"fmt"
)

func main() {
	type test struct {
		score  int
		member string
	}
	type t []test
	t1 := []test{
		{score: 901, member: "php1"},
		{score: 902, member: "php2"},
		{score: 903, member: "php3"},
		{score: 904, member: "php4"},
	}
	for _,i:=range t1{
		fmt.Println(i)
	}
}
