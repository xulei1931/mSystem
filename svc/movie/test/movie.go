package main

import (
	"regexp"
	"strconv"
	"strings"
)
func GetMovieDirectory(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*? rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func GetMovieWrite(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a href="/celebrity/.*?">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[0]
	}
	return main
	//return string(result[0][1])
}
func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 主演
func GetMovieMainActor(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
}

//类型==

func GetMovieType(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
}

// 制片国家movie_conutry
func GetMovieCountry(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">制片国家/地区:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 语言
func GetMovieLanguage(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">语言:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//<span property="v:initialReleaseDate" content="2018-11-09(中国大陆)">2018-11-09(中国大陆)</span>
// 上映时间 movie_on_line

func GetMovieOnLineTime(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:initialReleaseDate" content="(.*?)">`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
	//return string(result[0][1])
}

// 片长
//<span property="v:runtime" content="114">114分钟</span>
func GetMovieSpan(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 又名
func GetMovieNameAs(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">又名:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 图片

func GetMoviePic(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	//reg := regexp.MustCompile(`<a class="nbgnbg".*?><img src="(.*?)" title="点击看更多海报".*?></a>`)
	reg := regexp.MustCompile(`<img src="(.*?)" title="点击看更多海报"`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//<strong class="ll rating_num" property="v:average">6.8</strong>
// 评分
func GetMovieScores(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//desc

func GetMovieDesc(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:summary" class="">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}



// 获取电影id
func MovieId(url string) int {
	reg := regexp.MustCompile(`subject/(\d+)/.*`)
	result := reg.FindAllStringSubmatch(string(url), -1)

	if len(result) == 0 {
		return 0
	}
	int, _ := strconv.Atoi(result[0][1])
	return int

}