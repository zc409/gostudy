package main

import "fmt"

func main() {
	// 	age := 37
	// 	if age < 30 {
	// 		fmt.Println("young man")
	// 	} else if age < 50 {
	// 		fmt.Println("middle age man")
	// 	} else {
	// 		fmt.Println("old man")
	// 	}
	//作用域
	//age 变量只在if条件判断语句中生效
	if age := 38; age < 30 {
		fmt.Println("young man")
	} else if age < 50 {
		fmt.Println("middle age man")
	} else {
		fmt.Println("old man")
	}
	// fmt.Println("age") 在这里找不到age变量
}
