package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	X string `json:"name"`
	Y int    `json:"age"`
}

func main() {
	s1 := person{
		X: "小红",
		Y: 18,
	}
	js1, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("wrong because of %v", err)
		return
	}
	fmt.Printf("value:%v\n", string(js1))
	str := `{"name":"小明","age":18}`
	s2 := []byte(str)
	var s3 person
	json.Unmarshal(s2, &s3)
	fmt.Printf("value:%#v\n", s3)
	fmt.Println(s3)
	fmt.Println(s1)

}
