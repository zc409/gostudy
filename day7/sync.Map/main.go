package main

import (
	"fmt"
	"strconv"
	"sync"
)

// var m map[string]int

// func main(){
// 	wg:=sync.WaitGroup{}
// 	m=make(map[string]int,10)
// 	for i:=1;i<=21;i++{    //i设置超过20会报错：fatal error: concurrent map writes
// 		wg.Add(1)
// 		go func(n int){
// 			key:=strconv.Itoa(n)
// 			m[key]=n
// 			fmt.Printf("key=%s,value=%d\n",key,m[key])
// 			defer wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

func main() {
	m2 := sync.Map{} //初始化sync.Map类型
	wg := sync.WaitGroup{}
	for i := 1; i <= 100; i++ { //同时并发存储100个key
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, ok := m2.Load(key)
			if ok != true {
				fmt.Printf("can not load %s from m2!\n", key)
				return
			}
			fmt.Printf("key=%s,value=%d\n", key, value)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
