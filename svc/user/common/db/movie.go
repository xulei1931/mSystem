package db

import (
	"fmt"
	"user/common/entity"
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

// 获取正在上映的电影
func SelectTickingFilims(status int64) ([]entity.Film, error) {
	films := []entity.Film{}
	er := db.Where("is_ticking=?", status).Find(&films)
	if er.Error != nil { //
		return films, er.Error
	}
	return films, nil
}
