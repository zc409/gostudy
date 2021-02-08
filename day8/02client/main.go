package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1.与server端建立连接
	con, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("connet to server wrong,err:%v\n", err)
		return
	}
	//2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		//发送数据
		fmt.Print("请输入消息：")
		mes, _ := reader.ReadString('\n')
		mes = strings.TrimSpace(mes)
		mes = strings.ToLower(mes)
		if mes == "exit" {
			break
		} else {
			con.Write([]byte(mes))
		}

		//读数据
		//3.读取客户端发来消息
		tmp := make([]byte, 128)
		n, err := con.Read(tmp)
		if err != nil {
			fmt.Printf("read message wrong,err:%v\n", err)
			return
		}
		//4.打印客户端消息
		rmes := string(tmp[:n])
		fmt.Printf("%v\n", rmes)
	}
	con.Close()
}
