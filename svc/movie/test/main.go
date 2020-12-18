package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/goredis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"path"
	//"time"
	"regexp"
)

const (
	URL_QUEUE = "url_queue"
	URL_SER   = "url_set"
)

var (
	client goredis.Client
)
var db *gorm.DB

func init() {
	// 连接redis
	client.Addr = "127.0.0.1:6379"
	//InitDB()
}

// 定义一个初始化数据库的函数
func InitDB() error {
	// DSN:Data Source Name
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.SetLogger(Logger())
	fmt.Println(db.Error)
	return nil
}

//添加到队列
func PutQueue(url string) {
	client.Lpush(URL_QUEUE, []byte(url))
}

// 取出队列数据
func PopQueue() string {
	res, err := client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(res)

}

// 队列的长度
func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		return 0
	}
	return length
}

// 判断是否在集合中
func IsHave(url string) bool {
	has, err := client.Sismember(URL_SER, []byte(url))
	if err != nil {
		return false
	}
	return has
}
func GetMovieUrl(movieHtml string) []string {
	var movieurls []string
	// https://movie.douban.com/subject
	reg := regexp.MustCompile(`<a.*href="(https://movie.douban.com/subject/.*?)"`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	fmt.Println(result, "result")

	if len(result) == 0 {
		return movieurls
	}
	for _, v := range result {
		movieurls = append(movieurls, v[1])
	}
	return movieurls
}

type Subject struct {
	Subjects []Movie `json:"subjects"`
}

type Movie struct {
	TagId    int    `gorm:"column:tag_id"`
	MovieId  string `json:"id" gorm:"column:movie_id"`
	Title    string `json:"title" gorm:"column:title"`
	Url      string `json:"url" gorm:"column:url"`
	Rate     string `json:"rate" gorm:"column:rate"`
	PlayAble int    `json:"playable" gorm:"column:play_able"`
	IsNew    int    `json:"is_new" gorm:"column:is_new"`
	Cover    string `json:"cover" gorm:"column:cover"`
	CoverX   int64  `json:"cover_x" gorm:"column:cover_x"`
	CoverY   int64  `json:"cover_y" gorm:"column:cover_y"`
}

func (Movie) TableName() string {
	return "t_movies"
}

func main() {
	//c := cron.New()
	//spec := "*/10 * * * * ?"
	//c.AddFunc(spec, func() {
	//	now := time.Now()
	//	fmt.Println("cron running:", now.Minute(), now.Second())
	//})
	//c.Start()
	//select {}
	//
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbnane?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
//	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(Logger())
	defer db.Close()

	surl := "https://movie.douban.com/j/search_subjects?type=movie&tag=最新&sort=recommend&page_limit=20&page_start="
	//PutQueue(surl)
	var page int64=0
	for {
		url := surl + strconv.Itoa(int(page))
		fmt.Println(url)
		//os.Exit(0)
		//length := GetQueueLength()
		rsp := httplib.Get(url)

		subject, err := rsp.String()

		if err != nil {
			continue
		}
		fmt.Println(subject)
		bytes2 := []byte(subject)
		var files Subject
		err = json.Unmarshal(bytes2, &files)
		if err != nil {
			fmt.Println("json unmarshal error:",err)
		}
		fmt.Println(files)
		//  没数据自动退出
		if len(files.Subjects) ==0 {
			break
		}
		//fmt.Println(files)
		var u = new(Movie)
		for _, file := range files.Subjects {

		 er := db.Where("movie=?",file.MovieId).Find(&u)
			if er.Error == nil {
				// 存在了
             continue
			}

			MovieObject := Movie {
				TagId: 2,
				MovieId:  file.MovieId,
				Title:    file.Title,
				Url:      file.Title,
				Cover:    file.Cover,
				Rate:     file.Rate,
				PlayAble: file.PlayAble,
				IsNew:    file.IsNew,
				CoverX:   file.CoverX,
				CoverY:   file.CoverY,
			}
			//fmt.Println(&MovieObject)
			db.LogMode(true)
			db.Create(&MovieObject)
			//fmt.Println(err.Error)
			//fmt.Println()
		}
		atomic.AddInt64(&page, 20)
		time.Sleep(time.Second*1)
	}
	//defer db.Close()

}
func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}