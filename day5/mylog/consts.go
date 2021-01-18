package mylog

type loglevel int8

const (
	Info loglevel = iota
	Warn
	Error
	Fatal
)

//Mylog defination
type Mylog struct {
	level loglevel
}