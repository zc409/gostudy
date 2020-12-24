//1.求数组[1 3 5 7 8]所有元素的和
// package main

// import "fmt"

// func main() {
// 	sumre := 0
// 	aa := [...]int{1, 3, 5, 7, 8}
// 	for _, v := range aa {
// 		sumre += v
// 	}
// 	fmt.Println(sumre)
// }

//2.找出数组中和为指定值的两个元素的下标，比如数组[1 3 5 7 8]中找出和为8的两个元素的下标分别为（0，3）和（1，2）
package main

import "fmt"

func main() {
	ss := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i]+ss[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
}
