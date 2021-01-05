package main

import "fmt"

// func test(x int) int {
// 	if x == 1 {
// 		return 1
// 	} else {
// 		return x * test(x-1)

// 	}
// }

// func main() {
// 	fmt.Println(test(5))
// }
//}
func test(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	} else {
		return test(x-1) + test(x-2)
	}
}
func main() {
	fmt.Println(test(4))
}
