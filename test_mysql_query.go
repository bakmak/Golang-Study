/* package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB					//	定义一个全局对象db

type user struct {
	id int
	username string
	password string
}

func initDB() (err error) {		//	定义一个初始化数据库的函数
	dsn := "root:qwerty785621@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func queryRowDemo()  {						// 查询一条用户数据
	sqlStr := "select * from user_tbl where id = ?"
	// sqlStr := "select id, username, password from user_tbl where id=?"
	var u user
	err := db.QueryRow(sqlStr, 4).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed err: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%s\n", u.id, u.username, u.password)
}

func queryMultiRow()  {
	sqlStr := "select id, username, password from user_tbl where id > ?"
	rows, err := db.Query(sqlStr, 5)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return 
	}
	defer rows.Close()

	for rows.Next() {
		var u user 
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
		}
		fmt.Printf("id:%d username:%s password:%s\n", u.id, u.username, u.password)
	}
}

func main() {

	err := initDB()							//	调用输出化数据库的函数
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	// queryRowDemo()
	queryMultiRow()
} */