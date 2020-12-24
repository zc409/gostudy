package main

import "fmt"

// func main() {
// 	ss := [...]string{"北京", "上海", "广州"}
// 	for _, v := range ss {
// 		fmt.Println(v)
// 	}
// 	for i := 0; i < len(ss); i++ {
// 		fmt.Println(ss[i])
// 	}
// }
//三维数组
// func main() {
// 	var test [2][3][2]int
// 	test = [2][3][2]int{
// 		[3][2]int{
// 			[2]int{1, 2},
// 			[2]int{3, 4},
// 			[2]int{5, 6},
// 		},
// 		[3][2]int{
// 			[2]int{7, 8},
// 			[2]int{9, 10},
// 			[2]int{11, 12},
// 		},
// 	}
// 	fmt.Println(test)
// 	for i, v := range test {
// 		fmt.Printf("%d,%v\n", i, v)
// 	}
// }
// 二维数组
func main() {
	//	var test [3][2]int
	test := [...][2]int{
		[...]int{1, 2},
		[...]int{3, 4},
		[...]int{5, 6},
	}
	fmt.Println(test)
	for i, v := range test {
		fmt.Printf("%d,%v\n", i, v)
	}
}
