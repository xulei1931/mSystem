package db

import (
	"fmt"
	"movie/common/entity"
)

// 获取影片详情
func SelectFilmDetail(movieId int64) (*entity.Movie, error) {
	film := new(entity.Movie)
	fmt.Println("movie_id", movieId)
	er := db.Where("movie_id=?", movieId).Find(&film)
	if er.Error != nil { //
		return film, er.Error
	}
	return film, nil
}

// 获取正在上映的电影
func MovieLists(tag_id int64, limit int64, page int64) ([]entity.Movie, int32, error) {
	// 查询总数
	fileDB := db.Model(&entity.Movie{}).Where(&entity.Movie{TagId: int32(tag_id)})
	var count int32
	fileDB.Count(&count)

	films := []entity.Movie{}

	//er := db.Where("tag_id=?", tag_id).Find(&films)
	er := fileDB.Offset((page - 1) * limit).Limit(limit).Find(&films) //查询pageindex页的数据

	if er.Error != nil { //
		return films, count, er.Error
	}
	return films, count, nil
}
func GetTags(tag_id int64) ([]entity.Tags, int32, error) {
	var tags []entity.Tags
	var er error
	if tag_id <= 0 {
		// all
		er = db.Find(&tags).Error
	} else {
		er = db.Where("id=?", tag_id).Find(&tags).Error
	}
	if er != nil { //
		return tags, 0, er
	}
	return tags, int32(len(tags)), nil
}
