package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("connet to udp failed,err%v\n", err)
		return
	}
	defer con.Close()
	var data [1024]byte
	source := bufio.NewReader(os.Stdin)
	for {
		//发送数据
		fmt.Print("发送消息：")
		send, _ := source.ReadString('\n')
		_, err := con.Write([]byte(send))
		if err != nil {
			fmt.Printf("send mess failed,err:%v\n", err)
			continue
		}
		//接收消息
		n, _, err := con.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("receive mess failed,err:%v\n", err)
			continue
		}
		fmt.Printf("收到消息：%v", string(data[:n]))
	}
}
