package split_string

import (
	"strings"
)

//Split self strings.Split
func Split(source string, sep string) []string {
	res := make([]string, 0, strings.Count(source, sep)+1)
	run := true
	var index int
	for run {
		for i := 0; i < len(source); i++ {
			if string(source[i]) == sep {
				index = i
				run = true
				break
			} else {
				run = false
			}
		}
		if run == true {
			if index == 0 {
				if len(source) > 1 {
					source = source[1:]
				} else {
					break
				}
			} else {
				if index < len(source)-1 {
					res = append(res, source[:index])
					source = source[index+1:]
				} else {
					res = append(res, source[:index])
					break
				}
			}
		} else {
			res = append(res, source[:])
		}

	}
	return res
}

//Fib 斐波那契数列
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
