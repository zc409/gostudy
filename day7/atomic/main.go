package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg   sync.WaitGroup
	x    = int64(0)
	lock sync.Mutex
)

func add() {
	lock.Lock()
	x = x + 1
	lock.Unlock()
	defer wg.Done()
}

func add2() {
	atomic.AddInt64(&x, 1) //使用go的原子操作，不用自己添加锁。导入sync/atomic即可
	defer wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		//go add()   普通的加锁，性能没有atomic好，但是代码中执行时间没有看出来差别
		go add2()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("x=%v,costtime=%v", x, end.Sub(start))
}
