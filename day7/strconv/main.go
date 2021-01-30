package main

import (
	"fmt"
	"strconv"
)

func main() {
	//将字符串转化为int64
	str := "1000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	//Atoi:字符串转换成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	//Itoa:int转换成字符串
	a := 100
	str2 := strconv.Itoa(a)
	fmt.Printf("%T %v\n", str2, str2)

	//Sprintf将int转换成字符串
	str3:=fmt.Sprintf("%d",a)
	fmt.Printf("%T %v\n",str3,str3)
}
