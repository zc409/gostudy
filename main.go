package main

import "fmt"

type student struct {
	id   int
	name string
}

type mymap map[int]*student

var allstudent mymap

// func newStudent(id int,name string)*student{
// 	return &student{
// 		id: id,
// 		name: name,
// 	}
// }

func (a *mymap) addnew(id int, name string) {
	var newstudent = student{
		id:   id,
		name: name,
	}
	(*a)[id] = &newstudent
}

func main() {
	allstudent = make(mymap, 50)
	allstudent.addnew(1, "小明")
	fmt.Printf("type:%T,value:%v", allstudent, *allstudent[1])
}
