package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const dsn = "root:123456@tcp(127.0.0.1:3306)/blog"

func init() {
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	//DB.SetMaxOpenConns(20) // 设置连接池最大连接数
	//DB.SetMaxIdleConns(10)  // 设置连接池最大空闲连接数

	if err != nil {
		fmt.Println("数据库连接失败==>", err)
	}
	fmt.Println("数据库已连接！")
}

func InsertUser() {
	sql := `insert into user(id,name) values(?,?)`
	res, err := DB.Exec(sql, 2, "tm")
	fmt.Println(err)

	count, err := res.RowsAffected()
	fmt.Println(err)

	id, err := res.LastInsertId()
	fmt.Println(err)
	fmt.Println(count, id)
}

type User struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}

func QueryUserOne() {
	sql := `select id,name from user where id =1`
	u := User{}
	err := DB.Get(&u, sql)
	fmt.Println(err)
	fmt.Println(u)
}

func main() {
	//json.Marshal()
	//json.Unmarshal()
	//InsertUser()
	QueryUserOne()
}
