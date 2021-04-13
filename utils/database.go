package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName    = "root"
	password    = "123456"
	ip          = "bcore.top"
	port        = "3306"
	dbName      = "backend-admin"
)

var DB *sql.DB

func RegisterDb()  {
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, _ = sql.Open("mysql", path)
	//验证连接
	if errConn := DB.Ping(); errConn != nil{
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
	defer DB.Close()

	stmt, err := DB.Prepare("insert into user(name age) values(?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec("username", 18)
	if err != nil {
		fmt.Println(err)
	}
	// 获取新插入行的id
	fmt.Println(res.LastInsertId())
}

