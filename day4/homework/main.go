/* 函数版学生管理系统
写一个系统能够查看\新增\删除学生 */

package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

var allstudent map[int]*student

func judge(id int) bool {
	for k := range allstudent {
		if k == id {
			return false
		}
	}
	return true
}

func findall() {
	fmt.Println("学生系统的全部资料如下：")
	for k, v := range allstudent {
		fmt.Printf("学号：%d,姓名：%s\n", k, (*v).name)
	}
}

func addnew(id int, name string) {
	news := newStudent(id, name)
	allstudent[id] = news
}

func delold(id int) {
	delete(allstudent, id)
}

func main() {
	allstudent = make(map[int]*student, 50)
	for true {
		fmt.Println("欢迎进入学生管理系统！")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
		`)
		var choice int
		fmt.Print("请输入对应的操作号码：")
		fmt.Scanln(&choice)
		//fmt.Printf("您选择的操作号码是：%d\n", choice)
		switch choice {
		case 1:
			findall()
		case 2:
			var id int
			var name string
			fmt.Println("您好，您选择的是添加新的学生。")
			fmt.Print("请输入学生的学号：")
			fmt.Scanln(&id)
			fmt.Print("请输入学生的姓名：")
			fmt.Scanln(&name)
			if judge(id) {
				addnew(id, name)
				fmt.Println("学生已经新增成功！")
			} else {
				fmt.Println("您输入的学生学号已存在，没有进行任何操作！退回主界面")
			}
		case 3:
			var id int
			fmt.Println("您好，您选择的是删除学生。")
			fmt.Print("请输入学生的学号：")
			fmt.Scanln(&id)
			if !judge(id) {
				delold(id)
			} else {
				fmt.Println("您输入的学生学号不存在，没有进行任何操作！退回主界面")
			}
		case 4:
			os.Exit(1)

		}
	}

}
