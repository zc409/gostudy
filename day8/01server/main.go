package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

//等待连接，并打印通信数据
func pp(con net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		//3.读取客户端发来消息
		tmp := make([]byte, 128)
		n, err := con.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read message wrong,err:%v\n", err)
			return
		}
		//4.打印客户端消息
		rmes := string(tmp[:n])
		fmt.Printf("%v\n", rmes)

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
	}
}

func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("open tcp err:%v\n", err)
		return
	}
	for {
		//2.等待别人来跟我建立连接
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("receive message wrong,err:%v\n", err)
			return
		}
		go pp(con)
	}
}
