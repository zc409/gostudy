package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

var ch chan int
var done chan bool
var wg sync.WaitGroup

func addrand(c chan<- int) {
	x := rand.Intn(100)
	c <- x
	defer wg.Done()
}

func flog(c <-chan int,done chan<- bool) {
	f, err := os.OpenFile("D:/gogo/src/github.com/zc409/gostudy/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("openfile err:", err)
		done<-false
		return
	}
	for tmpdata := range c {
		fmt.Fprintln(f, tmpdata)
	}
	err=f.Close()
	if err!=nil{
		fmt.Printf("close file err:%v",err)
		done<-false
		return
	}
	done<-true
}

func main() {
	ch = make(chan int, 10)
	done=make(chan bool,1)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addrand(ch)
	}
	go flog(ch,done)
	go func() {
		wg.Wait()
		close(ch)
	}()
	judge,ok:=<-done
	if ok!=true{
		fmt.Println("err!chanel done is empty")
		return
	}
	if judge==true{
		fmt.Println("write log sucesess")
	}else{
		fmt.Println("write log wrong")
	}

}
