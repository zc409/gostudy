package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initdb() (err error) {
	sqn := `root:123456@tcp(192.168.56.101:3306)/gotest`
	db, err = sqlx.Connect("mysql", sqn)
	if err != nil {
		return
	}
	return
}

func sqlxselect() {
	sqlstr1 := `select id,name,salary from gotest where id=?;`
	var dat gotest
	//单条查询
	err := db.Get(&dat, sqlstr1, 2)
	if err != nil {
		fmt.Printf("get sqlstr1:%v failed,err:%v\n", sqlstr1, err)
		return
	}
	fmt.Printf("value is:%#v\n", dat)
	var datlist []gotest
	sqlstr2 := `select id,name,salary from gotest;`
	//多行查询
	err = db.Select(&datlist, sqlstr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("value is %#v\n", datlist)
}

func sqlxtrans() {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	sqlstr1 := "update gotest set salary=1000 where id=?;"
	_, err = tx.Exec(sqlstr1, 1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql:%v failed,err:%v\n", sqlstr1, err)
		return
	}
	sqlstr2 := "update gotest set salary=1500 where id=?;" //故意写错表名称那么第一步操作会回滚
	_, err = tx.Exec(sqlstr2, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql:%v failed,err:%v\n", sqlstr2, err)
		return
	}

	err = tx.Commit() //提交事务
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed,err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

type gotest struct {
	ID     int
	Name   string
	Salary int
}

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	sqlxselect()
	sqlxtrans()
}
