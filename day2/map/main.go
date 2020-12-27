package main

import "fmt"

func main() {
	var testmap map[string]int
	fmt.Println(testmap == nil)        //还没有初始化（没有在内存中开辟空间）
	testmap = make(map[string]int, 10) //要估算好该map容量，避免在程序运行期间在动态扩容
	testmap["上海"] = 1
	testmap["武汉"] = 27
	fmt.Printf("testmp=%v,len=%d\n", testmap, len(testmap))
	for x, y := range testmap {
		fmt.Printf("x=%v,y=%v\n", x, y)
	}
	fmt.Println(testmap["重庆"]) //如果不存在key，打印对应类型的零值
	s1, ok := testmap["武汉"]    //s1是value，ok是是否非空
	fmt.Println(s1, ok)
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(s1)
	}
	//map的遍历
	for key, value := range testmap {
		fmt.Println(key, value)
	}
	//删除
	delete(testmap, "上海")
	fmt.Println(testmap)
	delete(testmap, "不存在") //删除不存在的key不进行任何操作，不报错
	fmt.Println(testmap)
}
