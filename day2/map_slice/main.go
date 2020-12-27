package main

import "fmt"

func main() {
	//map和slice组合
	//元素类型为map的切片
	var s1 []map[int]string
	s1 = make([]map[int]string, 2, 5)
	s1[0] = make(map[int]string, 2)
	s1[0][2] = "test"
	s1[0][22] = "tt"
	fmt.Println(s1)
	//值为切片类型的map
	var testmap map[string][]int
	testmap = make(map[string][]int, 3)
	//	testmap["武汉"] = make([]int, 2, 2)
	//	testmap["武汉"][0] = 10
	testmap["武汉"] = []int{20, 30}
	fmt.Println(testmap)
}
