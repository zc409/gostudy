package main

//import "math"
import "fmt"

//浮点数
func main() {
	// math.MaxFloat32  //float32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T value is %v\n", f2, f2) //显示声明float32类型
	//  f1=f2          float32类型的值不能直接赋值给float64类型的变量
}
