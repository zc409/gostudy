package main

import "fmt"

func main() {
	//make 函数创造切片
	s1 := make([]int, 3, 5)
	fmt.Printf("s1 is %v,len is %d,cap is %d\n", s1, len(s1), cap(s1))
	var s2 []int
	s3 := []int{}
	s4 := make([]int, 0)
	fmt.Printf("s2 nil %v\n", s2 == nil) //len(s2)=0;cap(s2)=0;s2==nil
	fmt.Printf("s3 nil %v\n", s3 == nil) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Printf("s4 nil %v\n", s4 == nil) //len(s4)=0;cap(s4)=0;s4!=nil

	//切片的赋值
	s5 := []int{1, 2, 3}
	s6 := s5
	fmt.Printf("s5 is %v,s6 is %v\n", s5, s6)
	s5[1] = 200
	fmt.Printf("s5 is %v,s6 is %v\n", s5, s6)

	//切片的遍历
	for i := 0; i < len(s5); i++ {
		fmt.Println(s5[i])
	}

	for _, v := range s5 {
		println(v)
	}
}
