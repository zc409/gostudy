package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "123"
	b ,err:= strconv.ParseInt(a, 10, 64)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("type:%T,value:%v", b, b)

}
