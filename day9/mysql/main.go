package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //执行init()
)

//go连接mysql示例

func main() {
	//连接数据库
	dsn := "root:123456@tcp(192.168.56.101:3306)/gotest"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会效验用户名和密码是否正确,dsn格式不正确时候报错
	if err != nil {
		fmt.Printf("dsn:%s format is wrong,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("connect success!")
}
