package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "how do you do how are you"
	s2 := strings.Split(s1, " ")
	//	fmt.Printf("s2=%v,type(s2)=%T,len(s2)=%d,cap(s2)=%d\n", s2, s2, len(s2), cap(s2))
	countmap := make(map[string]int, 10)
	for _, x := range s2 {
		countmap[x]++
	}
	fmt.Println(s1)
	//对map按照key进行排序
	s3 := []string{}
	for key := range countmap {
		s3 = append(s3, key)
	}
	sort.Strings(s3)
	for _, key := range s3 {
		fmt.Printf("%s %d\n", key, countmap[key])
	}
}
