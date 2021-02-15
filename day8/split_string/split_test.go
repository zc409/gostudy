package split_string

import (
	"reflect"
	"testing"
)

func TestSplitmap(t *testing.T) { //测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	type test struct {
		str  string
		sep  string
		want []string
	}
	testmap := map[string]test{
		"case1": test{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case2": test{str: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
		"case3": test{str: "babcbdb", sep: "b", want: []string{"a", "c", "d"}},
	}
	for name, ts := range testmap {
		t.Run(name, func(t *testing.T) { //使用t.Run执行子测试
			got := Split(ts.str, ts.sep)
			if !reflect.DeepEqual(ts.want, got) { //因为slice不能直接比较，借助反射包中的方法比较
				t.Errorf("want:%v,got:%v", ts.want, got) //测试失败输出错误提示
			}
		})
	}
}

//go test -run=TestSplitmap/case1  只测试TestSplitmap函数中的case1
//go test -cover 测试代码覆盖率
// go test -cover -coverprofile=c.out  测试代码覆盖率图形化分析工具
// go tool cover -html=c.out
