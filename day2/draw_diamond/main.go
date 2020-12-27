package main

import "fmt"

func main() {
	lens := 5
	for x := 1; x <= lens; x++ {
		for y := x; y < lens; y++ {
			fmt.Printf(" ")
		}
		for z := 1; z < 2*x; z++ {
			fmt.Printf("*")
		}

		fmt.Println("")

	}
	for x := lens - 1; x >= 1; x-- {
		for y := x; y < lens; y++ {
			fmt.Printf(" ")
		}
		for z := 1; z < 2*x; z++ {
			fmt.Printf("*")
		}

		fmt.Println("")

	}

}
