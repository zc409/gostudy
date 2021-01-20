package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

func setstudent(s interface{}) {
	v := reflect.ValueOf(s).Elem()
	v.FieldByName("Name").Set(reflect.ValueOf("xiaoming"))
	v.FieldByName("Age").Set(reflect.ValueOf(18))
}

func main() {
	var stu student
	setstudent(&stu)
	fmt.Println(stu)
}
