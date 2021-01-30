package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func p(n int) {
// 	fmt.Println(n)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go p(i) //开启goroutine去执行函数p()
// 	}
// 	//main函数结束由main函数启动的goroutine函数也都结束，由于goroutine调用需要时间，因此需要等一下结束main函数
// 	time.Sleep(time.Second)
// }

//随机数
// func rd() {
// 	rand.Seed(time.Now().UnixNano())
// 	a := rand.Int()
// 	b := rand.Intn(10)
// 	fmt.Println(a, b)
// }

//导入sync模块，设置计数变量WaitGroup
// var wg sync.WaitGroup

// func main() {
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		//启动一个goroutine就登记+1
// 		//wg.Add(1)
// 		//匿名函数
// 		go func(n int) {
// 			//goroutine结束就登记-1
// 			defer wg.Done()
// 			fmt.Println(n)
// 		}(i)
// 	}
// 	//等待所有登记的goroutine都结束
// 	wg.Wait()
// }

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func(n int) {
// 			fmt.Println(n)
// 			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		defer wg.Done()
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		defer wg.Done()
		fmt.Println("B:", i)
	}
}

func main() {
	wg.Add(18)
	runtime.GOMAXPROCS(1) //默认跑满整个cpu线程数量
	go a()
	go b()
	wg.Wait()
}
