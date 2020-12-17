package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for start := 1; start <= i; start++ {
			if start == i {
				fmt.Printf("%d*%d=%d\n", start, i, start*i)
			} else {
				fmt.Printf("%d*%d=%d ", start, i, start*i)
			}
		}
	}
}
