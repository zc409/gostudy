package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type randint struct {
	x int
}

type result struct {
	rd  *randint
	sum int
}

var ch1 chan *randint
var ch2 chan *result
var wg sync.WaitGroup

func addch(ch chan<- *randint) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		tmpx := rand.Int()
		ch <- &randint{
			x: tmpx,
		}
		time.Sleep(time.Millisecond * 300)
	}
}

func chtoch(ch1 <-chan *randint, ch2 chan<- *result) {
	defer wg.Done()
	for i := range ch1 {
		tmpx := i.x
		res := 0
		for tmpx > 0 {
			res = res + tmpx%10
			tmpx = tmpx / 10
		}
		ch2 <- &result{
			rd:  i,
			sum: res,
		}
	}
}

func readch(c <-chan *result) {
	for i := range c {
		fmt.Println(i.rd.x, i.sum)
	}
}

func main() {
	ch1 = make(chan *randint, 100)
	ch2 = make(chan *result, 100)
	wg.Add(5)
	for y := 0; y < 2; y++ {
		go addch(ch1)
	}
	for i := 0; i < 3; i++ {
		go chtoch(ch1, ch2)
	}
	readch(ch2)
	wg.Wait()
}
