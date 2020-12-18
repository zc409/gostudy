package main

import "fmt"

// byte和rune类型
//go语言为了处理非ASCII码类型的字符，定义了新的rune类型

func main() {
	// s := "byte你好"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v(%c\n)", s[i], s[i])
	// }
	// for _, x := range s {
	// 	fmt.Printf("%v(%c)\n", x, x) //%c 字符
	// }
	//遍历方法无法打印utf8字符，需要下面range方法打印utf8字符
	s2 := "白萝卜"
	s3 := []rune(s2)
	fmt.Println(s3)
	fmt.Printf("%c\n", s3)
	fmt.Printf("%s\n", string(s3))
	s3[0] = '红'
	fmt.Println(string(s3))
	c1 := "红"
	c2 := '红'
	c3 := "h"
	c4 := 'h'
	fmt.Printf("c1:%T,c2:%T,c3:%T,c4:%T\n", c1, c2, c3, c4)

	//类型转换
	n1 := 10
	var n2 float64
	n2 = float64(n1)
	fmt.Println(n2)

}
