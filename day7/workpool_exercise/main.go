package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a chan int
var wg sync.WaitGroup

func tmpadd(n int) int {
	if n < 10 {
		return n
	}
	if n == 10 {
		return 1
	}
	return n%10 + tmpadd(n/10)
}

func randc(c chan<- int, n int) {
	for i := 1; i <= n; i++ {
		rand.Seed(time.Now().UnixNano() * int64(i))
		x := rand.Intn(100)
		c <- x
	}
	close(c)
	defer wg.Done()
}

func sumn(c <-chan int) {
	for {
		x, ok := <-c
		if !ok {
			break
		}
		fmt.Println(x, tmpadd(x))
	}
	defer wg.Done()
}

func main() {
	a = make(chan int)
	wg.Add(2)
	go randc(a, 20)
	go sumn(a)
	wg.Wait()
}
