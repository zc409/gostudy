package main

import (
	"fmt"
	"reflect"
	"strings"
)

type student struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

type stu struct {
	NAME string
	AGE  int
}

//s为指针类型变量，函数中要改变值必须传指针
func setstudent(s interface{}, m map[string]interface{}) {
	v := reflect.ValueOf(s).Elem()
	//求变量中字段数量
	for i := 0; i < v.NumField(); i++ {
		//查询tag名
		tname := v.Type().Field(i).Tag.Get("ini")
		//查询key名
		kname := v.Type().Field(i).Name
		//如果没有tag名，那么假定tag名为小写的key名
		if tname == "" {
			tname = strings.ToLower(v.Type().Field(i).Name)
		}
		//如果map中有tag名对应的字段，则赋值
		if value, ok := m[tname]; ok {
			//比较map中的字段类型和s变量中的字段类型是否相同
			if reflect.ValueOf(value).Type() == v.FieldByName(kname).Type() {
				v.FieldByName(kname).Set(reflect.ValueOf(value))
			}
		}
	}
}

func main() {
	var a student
	var b = map[string]interface{}{"name": "xiaoming", "age": 18}
	setstudent(&a, b)
	fmt.Printf("a=%#v\n", a)
	var c stu
	setstudent(&c, b)
	fmt.Printf("c=%#v\n", c)
}
