package main

import (
	"fmt"
)

type testtmp struct {
	str string
	sep string
}

type test []testtmp

func main() {
	tests :=make(test,3,10)
	tests = test{
		{str: "abc", sep: "b"},
		{str: "asc", sep: "s"},
		{str: "wer", sep: "e"},
	}
	for index, v := range tests {
		fmt.Printf("index=%v,value=%#v\n", index, v)
	}
	fmt.Printf("%#v\n", tests[0])
	tests = append(tests, testtmp{str: "test", sep: "test"})
	fmt.Printf("%v\n", tests[3].str)

}
