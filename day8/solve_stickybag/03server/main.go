package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/zc409/gostudy/day8/solve_stickybag/protocol"
)

func process(con net.Conn) {
	defer con.Close()
	newreader := bufio.NewReader(con)
	for {
		mess, err := protocol.Decode(newreader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("decode wrong,err:%v\n", err)
			break
		}
		fmt.Println(mess)
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("open listen wrong,err:%v\n", err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen wrong,err:%v\n", err)
			return
		}
		go process(con)
	}

}
