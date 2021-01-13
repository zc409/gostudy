package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("wrong because of %v", err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("wrong because of %v", err)
			return
		}
		fmt.Print(line)
	}
}
