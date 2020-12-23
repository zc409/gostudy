package main

import "fmt"

// func main() {
// 	for i := 0; i < 11; i++ {
// 		if i == 3 {
// 			goto testtag
// 		}
// 		fmt.Println(i)
// 	}
// testtag:
// 	fmt.Println("it's over")
// }

func main() {
breaktag:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break breaktag

			}
			fmt.Printf("%d,%d\n", i, j)
		}
	}
}

