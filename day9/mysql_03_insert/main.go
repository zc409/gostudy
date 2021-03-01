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

//插入数据函数
func inst() {
	// 1.写sql语句
	sqlstr := `insert into gotest(id,name,salary) values(?,?,?)`
	ret, err := db.Exec(sqlstr, 4, "t4", 400)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theid, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %v\n", theid)
}

//更新数据函数
func updateRowDemo() {
	sqlStr := "update gotest set salary=? where id=?;"
	ret, err := db.Exec(sqlStr, 600, 3)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //获取更新数据数量
	if err != nil {
		fmt.Printf("get rowsaffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success,affected rows:%d\n", n)
}

//删除数据函数
func deldata() {
	sqlstr := `delete from gotest where id=?`
	res, err := db.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("%v failed,err:%v\n", sqlstr, err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("del success,affected rows:%d\n", n)
}

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	//inst()
	// updateRowDemo()
	deldata()
}
