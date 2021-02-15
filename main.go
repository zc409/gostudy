package main

import (
	"fmt"
)

func main() {
	type test []struct {
		str string
		sep string
	}
	tests := test{
		{str: "abc", sep: "b"},
		{str: "asc", sep: "s"},
		{str: "wer", sep: "e"},
	}
	for index, v := range tests {
		fmt.Printf("index=%v,value=%#v\n", index, v)
	}
}
