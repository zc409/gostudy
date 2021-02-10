package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen udp port failed,err:%v\n", err)
		return
	}
	defer con.Close() //关闭连接写在err判断后面，如果打开端口失败则不会执行con.Close()，不会报错
	source := bufio.NewReader(os.Stdin)
	for {
		var data [1024]byte
		//读数据
		n, inaddr, err := con.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read udp message wrong,err%v\n", err)
			continue
		}
		fmt.Printf("收到发送发地址：%v,发来的消息是：%v", inaddr, string(data[:n]))
		//发送消息
		fmt.Print("发送消息：")
		reply, _ := source.ReadString('\n')
		_, err = con.WriteToUDP([]byte(reply), inaddr)
		if err != nil {
			fmt.Printf("reply mess wrong,err%v\n", err)
			continue
		}
	}

}
