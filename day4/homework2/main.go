package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

type mymap map[int]*student

var allstudent mymap

func (m *mymap) addnew(id int, name string) {
	var n = student{
		id:   id,
		name: name,
	}
	(*m)[id] = &n
}

func (m *mymap) judge(id int) bool {
	for k := range *m {
		if k == id {
			return false
		}
	}
	return true
}

func (m *mymap) showall() {
	fmt.Println("学生系统的全部学生信息如下：")
	for _, v := range *m {
		fmt.Printf("学号：%d,姓名：%s\n", (*v).id, (*v).name)
	}
}

func (m *mymap) delold(id int) {
	delete(*m, id)
}

func main() {
	allstudent = make(mymap, 50)
	for true {
		var choice int
		fmt.Println("欢迎光临学生系统！")
		fmt.Println(`
		1.查看系统中所有学生信息
		2.增加学生信息
		3.删除学生信息
		4.退出
		`)
		fmt.Print("请输入想要的操作对应的数字：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			allstudent.showall()
		case 2:
			var newid int
			var newname string
			fmt.Println("您选择的操作是添加学生信息。")
			fmt.Print("请输入新增学生的学号：")
			fmt.Scanln(&newid)
			fmt.Print("请输入新增学生的姓名：")
			fmt.Scanln(&newname)
			if allstudent.judge(newid) {
				allstudent.addnew(newid, newname)
			} else {
				fmt.Println("您输入的学生学号已存在请核对后在进行操作，没有进行任何操作，返回目录！")
			}
		case 3:
			var oldid int
			fmt.Println("您选择的操作是删除学生信息。")
			fmt.Print("请输入需要删除的学生的学号：")
			fmt.Scanln(&oldid)
			if !allstudent.judge(oldid) {
				allstudent.delold(oldid)
			} else {
				fmt.Println("您输入的学生学号不存在请核对后在进行操作，没有进行任何操作，返回目录！")
			}
		case 4:
			os.Exit(1)
		default:
			fmt.Println("您输入的操作有误，返回目录")
		}
	}
}
