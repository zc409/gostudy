package main

import (
	"fmt"
	"sync"
	"time"
)

var x = int(0)
var wg sync.WaitGroup
var rwlock sync.RWMutex

func write(n int) {
	rwlock.Lock()
	x = x + 1
	fmt.Printf("write id=%d,x=%d\n", n, x)
	//time.Sleep(time.Millisecond * 200)
	defer wg.Done()
	rwlock.Unlock()
}

func read(n int) {
	rwlock.RLock()
	time.Sleep(time.Second * 3)
	fmt.Printf("read id=%d,x=%d\n", n, x)
	defer wg.Done()
	rwlock.RUnlock()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i)
	}
	for x := 0; x < 100; x++ {
		wg.Add(1)
		go read(x)
	}
	wg.Wait()
}
