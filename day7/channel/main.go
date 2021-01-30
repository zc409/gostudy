package main

import (
	"fmt"
	"sync"
)

var a chan int
var wg sync.WaitGroup

func nobuf() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Printf("goroutine从通道a中取到了值：%v\n", x)
	}()
	a = make(chan int)
	defer close(a) //关闭通道
	a <- 10
	wg.Wait()
}

func buf() {
	a = make(chan int, 2)
	defer close(a) //关闭通道
	a <- 20
	x := <-a
	fmt.Printf("有缓存的通道可以直接取值，值为：%v\n", x)
}

func main() {
	nobuf() //goroutine从通道a中取到了值：10
	buf()   //有缓存的通道可以直接取值，值为：20
}
