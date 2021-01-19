package mylog

import (
	"strings"
)

type loglevel int8

const (
	Info loglevel = iota
	Warn
	Error
	Fatal
)

//Minelog interface
type Minelog interface {
	Info(contents string)
	Warn(contents string)
	Error(contents string)
	Fatal(contents string)
}

//Mylog defination
type Mylog struct {
	level loglevel
}

//Mylogfile defination
type Mylogfile struct {
	level    loglevel
	path     string
	filename string
	maxsize  int64
}

//Melog package two type log
func Melog(slevel string, logtype string) Minelog {
	logtype = strings.ToLower(logtype)
	switch logtype {
	case "c":
		return Newloger(slevel)
	case "f":
		return Newfileloger(slevel, "D:/gogo/src/github.com/zc409/gostudy/day5/homework_log/", "log.txt", 3*1024)
	default:
		return Newloger(slevel)
	}
}
