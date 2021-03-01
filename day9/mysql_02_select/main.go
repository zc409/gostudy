package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //定义全局变量

type gotest struct { //定义结构体存储数据库数据
	id     int
	name   string
	salary int
}

type restype []gotest //为存储多个记录，定义切片

var res restype //定义全局变量

func initdb() (err error) {
	dsn := "root:123456@tcp(192.168.56.101:3306)/gotest"
	db, err = sql.Open("mysql", dsn) //创建数据库的连接池db
	if err != nil {
		fmt.Printf("dsn format wrong,err:%v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("connect success!")
	db.SetMaxOpenConns(10) //设置连接池的最大连接
	return
}

func query(res restype, num int) {
	sqlstr := `select * from gotest where id=?;` //1.写查询sql语句,可以定义变量?
	for i := 1; i < num+1; i++ {
		rowobj := db.QueryRow(sqlstr, i)                                   //2.执行，从连接池里拿一个连接出来去数据库查询单条记录
		err := rowobj.Scan(&res[i-1].id, &res[i-1].name, &res[i-1].salary) //必须对rowobj对象调用scan方法，因为可以释放连接
		if err != nil {
			fmt.Printf("get data failed,err:%v\n", err)
			return
		}
	}
}

func querys(res restype) { //多行查询
	sqlstr := `select * from gotest where id > ?;` //查询语句
	rowsobj, err := db.Query(sqlstr, 0)            //拿一个连接来查询多条记录
	if err != nil {
		fmt.Printf("query %s failed,err:%v\n", sqlstr, err)
		return
	}
	defer rowsobj.Close() //一定不要忘记关闭连接
	for i := 1; rowsobj.Next(); i++ {
		err := rowsobj.Scan(&res[i-1].id, &res[i-1].name, &res[i-1].salary)
		if err != nil {
			fmt.Printf("get value failed,err:%v\n", err)
			return
		}
	}
}

func main() {
	num := 4
	res := make(restype, num, 10)
	err := initdb()
	if err != nil {
		fmt.Printf("init failed,err:%v\n", err)
		return
	}
	//	query(res, num)
	querys(res)
	fmt.Printf("res:%#v\n", res)
}
