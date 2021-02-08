package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("connect to server wrong,err:%v\n", err)
		return
	}
	mes := "this is test!"
	for i := 0; i < 10; i++ {
		con.Write([]byte(mes))
		time.Sleep(time.Millisecond)
	}
	defer con.Close()
}
