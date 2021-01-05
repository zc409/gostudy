package main

import (
	"fmt"
)

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func (d dog) wang() {
	fmt.Printf("%s汪汪汪\n", d.name)
}

func main() {
	s1 := newDog("xiaohei")
	s2 := newDog("xiaohua")
	s1.wang()
	s2.wang()
}
