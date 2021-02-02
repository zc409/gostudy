package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var ch1 chan int
var wg sync.WaitGroup
var once sync.Once

func add(c chan<- int) {
	x := rand.Intn(100)
	c <- x
	defer wg.Done()
}

func read(c chan int) {
	once.Do(func() { close(c) }) //读的时候只需要关闭一次通道
	for x := range c {
		fmt.Println(x)
	}
	defer wg.Done()
}

func main() {
	ch1 = make(chan int, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go add(ch1)
	}
	wg.Wait()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go read(ch1)
	}
	wg.Wait()
}
