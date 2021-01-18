package main

import (
	"fmt"
	"io"
	"os"
)

func insert(n int64, content string) {
	//移动用户输入的字节数
	seeksize := n
	//插入用户输入的内容
	contents := content
	//打开源文件
	srcf, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	//读取原始文件前部分内容
	tmpbg := make([]byte, seeksize)
	n2, err := srcf.Read(tmpbg)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	//打开临时文件并写入前半部分内容
	tmpf, err := os.OpenFile("./test.tmp", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	_, err = tmpf.Write(tmpbg[:n2])
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	//将用户输入内容写入临时文件
	_, err = tmpf.WriteString(contents)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if n2 == int(seeksize) {
		//读取源文件剩余内容并写入临时文件
		for {
			n3, err := srcf.Read(tmpbg)
			if err == io.EOF {
				break //此处不能用return，否则后面的操作都无法执行
			}
			if err != nil {
				fmt.Printf("err:%v\n", err)
				return
			}
			_, err = tmpf.Write(tmpbg[:n3])
			if err != nil {
				fmt.Printf("err:%v\n", err)
				return
			}
		}
		//srcf.Seek(seeksize,0) 指针已经移动不需要自定义移动指针
	}
	tmpf.Close()
	srcf.Close()
	os.Rename("./test.tmp", "./test.txt")
}

func main() {
	insert(6, "test")
}
