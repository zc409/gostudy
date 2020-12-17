package main

//字符串
import (
	"fmt"
	"strings"
)

func main() {
	//转义符
	path := "\"D:\\gogo\\src\\github.com\\zc409\\gostudy\\day1\""
	path2 := "'D:\\gogo\\src\\github.com\\zc409\\gostudy\\day1'"
	fmt.Println(path)
	fmt.Println(path2)

	s := "i'm ok"
	fmt.Println(s)

	//多行的字符串
	s2 := `
	这是一个
		测试
文档
	`
	fmt.Println(s2)

	s3 := `"D:\\gogo\\src\\github.com\\zc409\\gostudy\\day1\es"`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "实现"

	ss := name + world
	fmt.Println(ss)
	//Sprintf不打印内容，只能赋值
	ss1 := fmt.Sprintf("%s%s", name, world)
	//直接打印字符串拼接
	fmt.Printf("%s%s\n", name, world)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret[2])
	//s3 := `"D:\\gogo\\src\\github.com\\zc409\\gostudy\\day1\es"`,gogo索引是2不是1

	//判断包含
	fmt.Println(strings.Contains(ss, "实现"))
	fmt.Println(strings.Contains(ss, "不实现"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasPrefix(ss, "实现"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))
	//index
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "b"))
	fmt.Println(strings.LastIndex(s4, "b"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
