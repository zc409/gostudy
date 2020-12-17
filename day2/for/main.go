package main

import "fmt"

func main() {
	//基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//变种3无限循环
	// for {
	// 	fmt.Println("endless")
	// }

	//for range 循环
	s := "helo世界"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
