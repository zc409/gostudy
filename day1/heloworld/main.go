package main

import "fmt"

type kh666 int

const (
	kh kh666 = 1 << iota
	kha
	khb
	khc
	khd
)

const (
	khe kh666 = 16 >> iota
	khf
	khg
	khh
	khi
)

func main() {
	fmt.Println("NO.1:", kh)
	fmt.Println("NO.2:", kha)
	fmt.Println("NO.3:", khb)
	fmt.Println("NO.4:", khc)
	fmt.Println("NO.5:", khd)
	fmt.Println("NO.6:", khe)
	fmt.Println("NO.7:", khf)
	fmt.Println("NO.8:", khg)
	fmt.Println("NO.9:", khh)
	fmt.Println("NO.10:", khi)
}
