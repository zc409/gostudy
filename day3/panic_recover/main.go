package main

import "fmt"

func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover panic")
		}
	}()

	fmt.Println("b1")
	panic("there is bug")
	fmt.Println("b2")
}

func c() {
	fmt.Println("c")
}

func main() {
	a()
	b()
	c()
}
