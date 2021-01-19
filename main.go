package main

import (
	"fmt"
	"runtime"
)

func main() {
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("get line number wrong!\n")
		return
	}
	tmps := fmt.Sprintf("%v:%v", file, line)
	fmt.Print(tmps)
}
