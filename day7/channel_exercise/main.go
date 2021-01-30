package main

import (
	"fmt"
	"sync"
)

var c1 chan int
var c2 chan int
var wg sync.WaitGroup
var once sync.Once

// func toc1(c1t chan int) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		c1t <- i
// 	}
// 	close(c1t)
// }

// func c1t2(c1t chan int, c2t chan int) {
// 	defer wg.Done()
// 	for {
// 		x, ok := <-c1t
// 		if !ok {
// 			break
// 		}
// 		c2t <- x * x
// 	}
// 	once.Do(func() { close(c2t) })
// }

// func read(c chan int) {
// 	for x := range c {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	c1 = make(chan int, 10)
// 	c2 = make(chan int, 10)
// 	wg.Add(2)
// 	go toc1(c1)
// 	go c1t2(c1, c2)
// 	wg.Wait()
// 	read(c2)
// }
//单向通道
func toc1(c1t chan<- int) { //receiveonly
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		c1t <- i
	}
	close(c1t)
}

func c1t2(c1t <-chan int, c2t chan<- int) { //c1t receive-only,c2t send-only
	defer wg.Done()
	for {
		x, ok := <-c1t
		if !ok {
			break
		}
		c2t <- x * x
	}
	once.Do(func() { close(c2t) })
}

func read(c <-chan int) { //receive-only
	for x := range c {
		fmt.Println(x)
	}
}

func main() {
	c1 = make(chan int, 10)
	c2 = make(chan int, 10)
	wg.Add(2)
	go toc1(c1)
	go c1t2(c1, c2)
	wg.Wait()
	read(c2)
}
