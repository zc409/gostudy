package main

import "fmt"

/* func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}  */

/* func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
} */

func main() {
	/* 	n := 3
	   	switch n {
	   	case 1:
	   		{
	   			fmt.Println("first")
	   		}
	   	case 2:
	   		{
	   			fmt.Println("second")
	   		}
	   	case 3:
	   		{
	   			fmt.Println("third")
	   		}
	   	case 4:
	   		{
	   			fmt.Println("forth")
	   		}
	   	case 5:
	   		{
	   			fmt.Println("fifth")
	   		}
	   	} */

	switch n := 3; n {
	case 1:
		{
			fmt.Println("first")
		}
	case 2:
		{
			fmt.Println("second")
		}
	case 3:
		{
			fmt.Println("third")
		}
	case 4:
		{
			fmt.Println("forth")
		}
	case 5:
		{
			fmt.Println("fifth")
		}
	default:
		{
			fmt.Println("other")
		}
	}
}
