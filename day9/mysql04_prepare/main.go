package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type gotest struct { //定义结构体存储数据库数据
	id     int
	name   string
	salary int
}

type restype []gotest //为存储多个记录，定义切片

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

func prepareInsert() {
	sqlStr := `insert into gotest(id,name,salary) values(?,?,?);`
	stmt, err := db.Prepare(sqlStr) //把sql语句先发给mysql预处理一下
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close() //不要忘记关闭连接
	// 后续只需要拿到stmt去执行操作
	data := restype{
		{id: 1, name: "t1", salary: 1},
		{id: 2, name: "t2", salary: 2},
		{id: 3, name: "t3", salary: 3},
	}
	for k := range data {
		_, err := stmt.Exec(data[k].id, data[k].name, data[k].salary) //后续自用传值
		if err != nil {
			fmt.Printf("exec failed,err:%v\n", err)
			return
		}
		fmt.Printf("insert key:%v,success!\n", k)
	}
}

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	prepareInsert()
}
