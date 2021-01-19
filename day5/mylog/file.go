package mylog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

//Newfileloger create Mylogfile type
func Newfileloger(l string, path string, filename string, maxsize int64) Mylogfile {
	dl := strings.ToLower(l)
	var le loglevel
	switch dl {
	case "info":
		le = Info
	case "warn":
		le = Warn
	case "error":
		le = Error
	case "fatal":
		le = Fatal
	}
	return Mylogfile{
		level:    le,
		path:     path,
		filename: filename,
		maxsize:  maxsize,
	}
}

//比较日志类型中的级别和输入级别，如果小与等于输入级别则返回true
func (m Mylogfile) compare(l string) bool {
	llow := strings.ToLower(l)
	var le loglevel
	switch llow {
	case "info":
		le = Info
	case "warn":
		le = Warn
	case "error":
		le = Error
	case "fatal":
		le = Fatal
	}
	return m.level <= le
}

func (m Mylogfile) printlog(slevel string, content string) {
	if m.compare(slevel) {
		timenow := time.Now().Format("2006-01-02 15:04:05")
		fname := path.Join(m.path, m.filename)
		fileobj, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open file err,err:%v\n", err)
			return
		}
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			fmt.Printf("get line number err!\n")
			return
		}
		linems := fmt.Sprintf("%v:%v", file, line)
		_, err = fmt.Fprintf(fileobj, "%v [%s] %s %s\n", timenow, slevel, linems, content)
		if err != nil {
			fmt.Printf("write file err,err:%v\n", err)
			return
		}
		fileobj.Close()
		fstat, err := os.Stat(fname)
		if err != nil {
			fmt.Printf("get file info err:%v\n", err)
			return
		}
		timenow2 := time.Now().Format("20060102_150405")
		newfname := fmt.Sprintf("%v_%v", fname, timenow2)
		if fstat.Size() >= m.maxsize {
			os.Rename(fname, newfname)
		}
	}
}

//Info log level
func (m Mylogfile) Info(contents string) {
	m.printlog("info", contents)
}

//Warn log level
func (m Mylogfile) Warn(contents string) {
	m.printlog("warn", contents)
}

//Error log level
func (m Mylogfile) Error(contents string) {
	m.printlog("error", contents)
}

//Fatal log level
func (m Mylogfile) Fatal(contents string) {
	m.printlog("fatal", contents)
}

//Off ,turn off log
func (m Mylogfile) Off(interface{}) {}
