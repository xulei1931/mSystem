package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
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
// UserInfo 用户信息
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
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

	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "七米", "男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	//// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.Where("id=?","1").Find(&u)
	fmt.Printf("%#v\n", u)
}