package db

import (
	"fmt"
	"mSystem/svc/common/entity"
)

// 获取影片详情
func SelectFilmDetail(movieId int64) (*entity.Film, error) {
	film := new(entity.Film)
	fmt.Println("movie_id", movieId)
	er := db.Where("movie_id=?", movieId).Find(&film)
	if er.Error != nil { //
		return film, er.Error
	}
	return film, nil
}
