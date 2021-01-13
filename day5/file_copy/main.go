package main

import (
	"fmt"
	"io"
	"os"
)

func copyfile() {
	f, err := os.Open("./main.go") //获取读文件对象
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer f.Close()

	dts, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) //获取写文件对象
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer dts.Close()

	n, err := io.Copy(dts, f) //copy读对象到写对象
	if err != nil {
		fmt.Printf("copy error,err:%v\n", err)
		return
	}
	fmt.Printf("success,write %v", n)
}

func main() {
	copyfile()
}
