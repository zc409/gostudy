package mylog

import (
	"fmt"
	"strings"
	"time"
)



//Newloger create Mylog type
func Newloger(l string) Mylog {
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
	return Mylog{
		level: le,
	}
}

func (m Mylog) compare(l string) bool {
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

func (m Mylog) printlog(slevel string, content string) {
	if m.compare(slevel) {
		timenow := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%v [%s] %s\n", timenow, slevel, content)
	}
}

//Info log level
func (m Mylog) Info(contents string) {
	m.printlog("info", contents)
}

//Warn log level
func (m Mylog) Warn(contents string) {
	m.printlog("warn", contents)
}

//Error log level
func (m Mylog) Error(contents string) {
	m.printlog("error", contents)
}

//Fatal log level
func (m Mylog) Fatal(contents string) {
	m.printlog("fatal", contents)
}

//Off ,turn off log
func (m Mylog) Off(interface{}) {}
