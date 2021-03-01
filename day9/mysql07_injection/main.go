package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type gotest struct {
	ID     int
	Name   string
	Salary int
}

func initdb() (err error) {
	sqn := `root:123456@tcp(192.168.56.101:3306)/gotest`
	db, err = sqlx.Connect("mysql", sqn)
	if err != nil {
		return
	}
	return
}

// sql注入
func sqlinjectiondemo(name string) {
	sqlstr := fmt.Sprintf("select id,name,salary from gotest where name='%s'", name)
	fmt.Printf("sql:%s\n", sqlstr)

	var datlist []gotest
	//多行查询
	err := db.Select(&datlist, sqlstr)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	for _, v := range datlist {
		fmt.Printf("%#v\n", v)
	}

}

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	sqlinjectiondemo("xxx' or 1=1#")
	// sqlinjectiondemo("xiaoming")
}
