package main

import (
	"fmt"
	"net"

	"github.com/zc409/gostudy/day8/solve_stickybag/protocol"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("connect to server wrong,err:%v\n", err)
		return
	}
	mes := "this is test!"
	mess, err := protocol.Encode(mes) //调用自定义protocol包编码消息
	if err != nil {
		fmt.Printf("encode wrong,err:%v\n", err)
		return
	}
	for i := 0; i < 10; i++ {
		con.Write(mess)
	}
	defer con.Close()
}
