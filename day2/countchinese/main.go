package main

import "fmt"

func main() {
	tmp1 := "this 理想 is 现实 猜猜多少个"
	tmp2 := []rune(tmp1)
	fmt.Printf("汉字个数是:%d", (len(tmp1)-len(tmp2))/2)
}
