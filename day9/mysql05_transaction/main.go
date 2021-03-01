package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initdb() (err error) {
	sqn := `root:123456@tcp(192.168.56.101:3306)/gotest`
	db, err = sql.Open("mysql", sqn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil { //如果有错误且事务开始有返回值
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed,err:%v\n", err)
		return
	}
	sqlstr1 := "update gotest set salary=2000 where id=?;"
	_, err = tx.Exec(sqlstr1, 1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql:%v failed,err:%v\n", sqlstr1, err)
		return
	}
	sqlstr2 := "update gotest set salary=1000 where id=?;" //故意写错表名称那么第一步操作会回滚
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

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	transactionDemo()
}
