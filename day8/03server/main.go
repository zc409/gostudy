package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(con net.Conn) {
	defer con.Close()
	for {
		newreader := bufio.NewReader(con)
		var buf [1024]byte
		n, err := newreader.Read(buf[:])
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		if err != nil {
			fmt.Printf("read mesage wrong,err:%v\n", err)
			break
		}
		mess := string(buf[:n])
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
