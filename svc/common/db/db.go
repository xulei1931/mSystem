package db

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 定义一个全局对象db
var db *gorm.DB
var redisClient *redis.Client

func Init() {
	initDB()
	initRedis()
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	db, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	//defer db.Close()
	return nil
}
func initRedis() {
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
func Close(){
	db.Close()
}