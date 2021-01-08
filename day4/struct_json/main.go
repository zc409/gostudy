package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//序列化
	p1 := person{
		Name: "小明",
		Age:  18,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err %v", err)
		return
	}
	fmt.Println(string(b))
	fmt.Printf("%#v\n", string(b))

	//反序列化
	str := []byte(`{"name":"小红","age":18}`)
	var stud person
	json.Unmarshal(str, &stud) //函数传值必须要带入内存地址
	fmt.Println(stud)
	fmt.Printf("%#v\n", stud)
	fmt.Println(p1.Name)
}
