package main

import (
	"fmt"
	"reflect"
)

func setvalue(x interface{}) {
	//v := reflect.ValueOf(x)
	// if v.Elem().Kind()==reflect.Int{
	// 	v.Elem().SetInt(365)
	v := reflect.ValueOf(x).Elem()
	if v.Kind() == reflect.Int {
		v.SetInt(365)
	}
}

func main() {
	a := 1
	setvalue(&a)
	fmt.Printf("type:%T,value:%v", a, a)
}
