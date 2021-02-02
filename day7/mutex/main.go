package main

import (
	"fmt"
	"sync"
)

var x = int(0)
var wg sync.WaitGroup
var lock sync.Mutex

func add1(n *int) {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		*n = *n + 1
		lock.Unlock() //解锁
	}
	defer wg.Done()
}

func main() {
	wg.Add(2)
	go add1(&x)
	go add1(&x)
	wg.Wait()
	// add1(&x)
	// add1(&x)
	fmt.Println(x)
}
