package main

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/goredis"
	"fmt"
	//"time"
	"regexp"
	"encoding/json"
)

const (
	URL_QUEUE = "url_queue"
	URL_SER   = "url_set"
)

var (
	client goredis.Client
)
func init(){
	// 连接redis
	client.Addr = "127.0.0.1:6379"
}
type MovieInfo struct {
	Id                   int64
	Movie_id             int
	Movie_name           string
	Movie_director       string
	Movie_writer         string
	Movie_conutry        string
	Movie_language       string
	Movie_type           string
	Movie_image          string
	Movie_main_charactor string
	Movie_on_line        string
	Movie_span           string
	Movie_grade          string
	remark               string
	_create_time         string
	Movie_name_as        string
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
	fmt.Println(result,"result")

	if len(result) == 0 {
		return movieurls
	}
	for _, v := range result {
		movieurls = append(movieurls, v[1])
	}
	return movieurls
}
type Subject struct {

	subjects []MovieObject
}
type MovieObject struct {
	rate float64  `json:"rate"`
	cover_x int64 `json:"cover_x"`
	title string `json:"title"`
	url string `json:"url"`
	playable bool `json:"playable"`
	cover string `json:"cover"`
	id int64 `json:"id"`
	is_new bool `json:"is_new"`
	
}
func main()  {
	surl := `https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=0`
	PutQueue(surl)
	for {
		length := GetQueueLength()
		fmt.Println(length)
		if length == 0 {
			break
		}
		surl = PopQueue()
		if IsHave(surl) {
			continue
		}
		rsp := httplib.Get(surl)


		subject, err := rsp.String()
		fmt.Println(subject)
		if err != nil {
			continue
		}

		bytes2 := []byte(subject)
		var files Subject
		if json.Unmarshal(bytes2, &files) == nil {
			//fmt.Println("json.Unmarshal 解码结果: ", person2.Name, person2.Age)
		}

		fmt.Println(files)
		// 获取url
		//var movieInfo MovieInfo
		//movieInfo.Movie_name = GetMovieName(html)
		//if movieInfo.Movie_name != "" {
		//	// 爬取电影
		//	movieInfo.Id = 0
		//	movieInfo.Movie_id = MovieId(surl)
		//	movieInfo.Movie_name = GetMovieName(html)
		//	movieInfo.Movie_director = GetMovieDirectory(html)
		//	movieInfo.Movie_writer = GetMovieWrite(html)
		//	movieInfo.Movie_conutry = GetMovieCountry(html)
		//	movieInfo.Movie_language = GetMovieLanguage(html)
		//	movieInfo.Movie_type = GetMovieType(html)
		//	movieInfo.Movie_image = GetMoviePic(html)
		//	movieInfo.Movie_main_charactor = GetMovieMainActor(html)
		//	movieInfo.Movie_on_line = GetMovieOnLineTime(html)
		//	movieInfo.Movie_span = GetMovieSpan(html)
		//	movieInfo.Movie_grade = GetMovieScores(html)
		//	movieInfo.Movie_name_as = GetMovieNameAs(html)
		//	fmt.Println(movieInfo)
		//	time.Sleep(time.Second)
		//}
	}
}