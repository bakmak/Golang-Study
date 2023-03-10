/* package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
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

func inserDate()  {
	sqlStr := "insert into user_tbl(username, password) values(?,?)"
	ret, err := db.Exec(sqlStr, "张三", "zs123")
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}

func insert2(username string, password string) {
	sqlStr := "insert into user_tbl(username, password) values(?,?)"
	r, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", r)
	}
	i, err := r.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", i)
}


func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	}else{
		fmt.Printf("初始化成功\n")
	}

	// inserDate()
	insert2("李四","lisi123")
}
 */