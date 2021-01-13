package main

import "fmt"

//类型断言
func assigs(a interface{}) {
	fmt.Printf("%T\n", a)
	x, y := a.(string)
	fmt.Printf("x:%v,y:%v\n", x, y)
}

func assigs2(a interface{}) { //switch
	switch a.(type) {
	case string:
		fmt.Println(a, "this is string")
	case bool:
		fmt.Println(a, "this is bool")
	case int:
		fmt.Println(a, "this is int")
	}
}

func main() {
	assigs(100)
	assigs("heloworld")
	assigs2("helo")
	assigs2(100)
	assigs2(true)
}
