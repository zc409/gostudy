package main

import "fmt"
//标识符：变量名 函数名 类型名 方法名
//go语言中如果标识符首字母是大写的，就表示对外部包可见

// Dog 首字母大写需要添加注释
type Dog struct {
	name   string
	gender string
}

//构造函数
func newDog(name string, gender string) dog {
	return dog{
		name:   name,
		gender: gender,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用于类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s哇哇哇，性别是：%s\n", d.name, d.gender)
}

func main() {
	s1 := newDog("xiaohei", "male")
	s2 := newDog("xiaohua", "female")
	s1.wang()
	s2.wang()
}
