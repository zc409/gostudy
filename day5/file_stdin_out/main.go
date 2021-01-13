package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	stdin, err := reader.ReadString('\n')
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	writer := bufio.NewWriter(f)
	writer.WriteString(stdin)
	writer.Flush()
	fmt.Fprint(f, stdin)         //往文件对象中写
	fmt.Fprint(os.Stdout, stdin) //往标准输出中写
}
