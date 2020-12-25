package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	// s1[3] = "广州" //错误的写法 会导致编译错误：索引越界
	// fmt.Pringln(s1)
	s1 = append(s1, "武汉")
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := [...]string{"重庆", "成都"}
	s1 = append(s1, ss[:]...)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
}
