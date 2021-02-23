package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan,value:%v\n", v)
		default:

		}
	}
}

func main() {
	var isCpuPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on") //是否开启cpupprof的标志位
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on") //是否开启内存pprof的标志位
	flag.Parse()

	if isCpuPprof {
		file, err := os.Create("./cpu.pprof") //在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(10 * time.Second)
	if isMemPprof {
		file2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed,err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file2)
		defer file2.Close()
	}
}

// D:\gogo\src\github.com\zc409\gostudy\day8\pprof_demo>go tool pprof cpu.pprof
// Type: cpu
// Time: Feb 22, 2021 at 2:53pm (CST)
// Duration: 10.17s, Total samples = 1.29mins (762.26%)
// Entering interactive mode (type "help" for commands, "o" for options)
// (pprof) top
// Showing nodes accounting for 77.20s, 99.60% of 77.51s total
// Dropped 21 nodes (cum <= 0.39s)
//       flat  flat%   sum%        cum   cum%
//     35.58s 45.90% 45.90%     63.67s 82.14%  runtime.selectnbrecv
//     28.08s 36.23% 82.13%     28.08s 36.23%  runtime.chanrecv
//     13.54s 17.47% 99.60%     77.30s 99.73%  main.logicCode
// (pprof) list logicCode
// Total: 1.29mins
// ROUTINE ======================== main.logicCode in D:\gogo\src\github.com\zc409\gostudy\day8\pprof_demo\main.go
//     13.54s   1.29mins (flat, cum) 99.73% of Total
//          .          .     10:
//          .          .     11:func logicCode() {
//          .          .     12:   var c chan int
//          .          .     13:   for {
//          .          .     14:           select {
//     13.54s   1.29mins     15:           case v := <-c:
//          .          .     16:                   fmt.Printf("recv from chan,value:%v\n", v)
//          .          .     17:           default:
//          .          .     18:
//          .          .     19:           }
//          .          .     20:   }
// (pprof) quit
