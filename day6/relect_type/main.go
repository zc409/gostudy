package main

import (
	"fmt"
	"reflect"
)

type student struct{}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func main() {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b) //type:int64
	var c = student{}
	//TypeOf 获取的时type name，想获取底层类型需要后去type的kind，如下
	fmt.Printf("type name:%v,type kind:%v\n", reflect.TypeOf(c).Name(), reflect.TypeOf(c).Kind()) //type name:student,type kind:struct
}
