package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	CreateAt string `json:"create_at" db:"create_at"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
}
// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True"
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
func main()  {

	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	user := User{}
	sqlStr := "SELECT * FROM user WHERE `email` = ?"
	e := db.QueryRow(sqlStr,"342591255@qq.com").Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateAt,&user.Email,&user.Phone)
	if e != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println(user)
}