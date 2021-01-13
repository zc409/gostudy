package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func common() {
	f, err := os.Open("./main.go")
	if err != nil {
		//fmt.Printf("there is problem:%v\n", err)
		return
	}
	defer f.Close()
	tmp := make([]byte, 20)
	for {
		n, err := f.Read(tmp)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("wrong because %v\n", err)
			return
		}
		//fmt.Printf("n:%v\n", n)
		fmt.Print(string(tmp[:n]))
	}
}

func readbyioutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read wrong,err:%v\n", err)
		return
	}
	fmt.Print(string(ret))
	fmt.Println("this is read by readbyioutil")
}

func main() {
	//	common()
	readbyioutil()
}
