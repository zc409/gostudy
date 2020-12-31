package main

//var num int64 = 10

// func testGlobalvar() {
// 	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
// 	//函数中查找变量的顺序
// 	//1.现在函数内部查找
// 	//2.找不到就往函数的外面查找，一直找到全局
// }
// func main() {
// 	testGlobalvar() //num=10
// }

// func testLocalVar() {
// 	//定义一个函数局部变量x，仅在该函数内生效
// 	var x int64 = 100
// 	fmt.Printf("x=%d\n", x)
// }

// func main() {
// 	testLocalVar()
// 	fmt.Println(x) //此时无法使用变量x
// }

func main() {
	//语句块作用域
	// if i := 0; i < 18 {
	// 	fmt.Println("乖乖上学")
	// }
	// fmt.Println(i)  //无法使用变量i
	// for j:=0;j<5;j++{
	// 	fmt.Println(j)
	// }
	// fmt.Print(j)  //无法使用变量j
}
