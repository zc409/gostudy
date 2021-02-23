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

func BenchmarkSplit(b *testing.B) {
	// time.Sleep(time.Second)
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e:f:g:h", ":")
	}
}

// go test -bench=Split
// goos: windows
// goarch: amd64
// pkg: github.com/zc409/gostudy/day8/split_string
// BenchmarkSplit-8         2724954               476 ns/op
// PASS
// ok      github.com/zc409/gostudy/day8/split_string      1.944s
// D:\gogo\src\github.com\zc409\gostudy\day8\split_string>go test -bench=Split -benchmem
// goos: windows
// goarch: amd64
// pkg: github.com/zc409/gostudy/day8/split_string
// BenchmarkSplit-8         4609963               264 ns/op             128 B/op          1 allocs/op
// PASS
// ok      github.com/zc409/gostudy/day8/split_string      1.688s

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

// go test -bench=Fib
// goos: windows
// goarch: amd64
// pkg: github.com/zc409/gostudy/day8/split_string
// BenchmarkFib1-8         469979841                2.33 ns/op
// BenchmarkFib2-8         184654132                6.45 ns/op
// BenchmarkFib3-8         100000000               11.2 ns/op
// BenchmarkFib10-8         2800629               364 ns/op
// BenchmarkFib20-8           29588             41642 ns/op
// BenchmarkFib40-8               2         604488000 ns/op
// PASS
// ok      github.com/zc409/gostudy/day8/split_string      9.527s
// -benchtime 自定义测试时间（默认为1s）
// D:\gogo\src\github.com\zc409\gostudy\day8\split_string>go test -bench=Fib40 -benchtime=20s
// goos: windows
// goarch: amd64
// pkg: github.com/zc409/gostudy/day8/split_string
// BenchmarkFib40-8              38         597037187 ns/op
// PASS
// ok      github.com/zc409/gostudy/day8/split_string      23.543s

func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1) //设置使用的cpu数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}

// D:\gogo\src\github.com\zc409\gostudy\day8\split_string>go test -bench=Split
// goos: windows
// goarch: amd64
// pkg: github.com/zc409/gostudy/day8/split_string
// BenchmarkSplit-8                 4406806               264 ns/op
// BenchmarkSplitParallel-8        31020656                40.5 ns/op
// PASS
// ok      github.com/zc409/gostudy/day8/split_string      2.935s
