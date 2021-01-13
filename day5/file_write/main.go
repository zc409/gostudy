package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write1() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Printf("there is problem,err:%v", err)
		return
	}
	//write
	f.Write([]byte("你好，很高兴见到你！\n"))

	//writestring
	f.WriteString("这是字符串！\n")
}

func write2() {
	f, err := os.OpenFile("./test2.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		writer.WriteString("循环i次\n")
	}
	writer.Flush()
}

func write3() {
	str := `
这是一个测试
社么
这个时测试？
	哈哈
	`
	err := ioutil.WriteFile("./test3.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
}

func main() {
	//write1()
	//write2()
	write3()
}
