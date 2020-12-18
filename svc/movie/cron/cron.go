package cron

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/jinzhu/gorm"
	"mSystem/svc/common/entity"
	"strconv"
	"sync/atomic"
	"time"
)

// 定时任务更新电影
const url = "https://movie.douban.com/j/search_subjects?type=movie&sort=recommend&page_limit=20"

func InitCron(db *gorm.DB){
	//c := cron.New()
	////具体定时函数
	//spec := "*/100 * * * * ?"
	//c.AddFunc(spec, func() {
	//	now := time.Now()
	//	SyncMocie(db)
	//	fmt.Println("cron running:", now.Minute(), now.Second())
	//})
	//c.Start()
	SyncMocie(db)
}
func SyncMocie(db *gorm.DB){
	// 获取tags
	tags := []entity.Tags{}
	err :=db.Find(&tags)
	if err.Error ==nil{
       for _,tag := range tags{
		   UpdateMovie(db,tag.CatName,tag.Id)
	   }
	}
}
func UpdateMovie(db *gorm.DB,tag_name string,tag_id int){
	var page int64=0
	for {
		url := url + "&tag=" + tag_name + "&page_start=" + strconv.Itoa(int(page))
		fmt.Println(url)

		rsp := httplib.Get(url)

		subject, err := rsp.String()

		if err != nil {
			continue
		}
		bytes2 := []byte(subject)
		var files entity.Subject
		err = json.Unmarshal(bytes2, &files)
		if err != nil {
			fmt.Println("json.Unmarshal 解码错误: ", err)
		}
		//  没数据自动退出
		if len(files.Subjects) == 0 {
			fmt.Println(tag_name, "sync success!!!")
			break
		}

		var u = new(entity.Movie)
		for _, file := range files.Subjects {
			er := db.Where("movie_id=?",file.MovieId).Find(&u)
			if er.Error == nil {
				// 存在了
				continue
			}
			MovieObject := entity.Movie {
				TagId: tag_id,
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
			//fmt.Println(MovieObject)
			db.Create(&MovieObject)
		}
		atomic.AddInt64(&page, 20)
		time.Sleep(time.Millisecond*100)
	}
}