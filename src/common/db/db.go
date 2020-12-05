package db

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)
// 定义一个全局对象db
var db *sql.DB
var redisClient *redis.Client

func Init()  {
	initDB()
	initRedis()
}
// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:1234@tcp(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True"
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
func initRedis(){
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis 连接失败。。。。。")
	}
	fmt.Println(pong, "redis 连接成功！！！")
}