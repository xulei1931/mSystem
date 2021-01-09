package entity

type Subject struct {
	Subjects []Movie `json:"subjects"`
}

type Movie struct {
	TagId    int32    `gorm:"column:tag_id"`
	MovieId  int64 `json:"id" gorm:"column:movie_id"`
	Title    string `json:"title" gorm:"column:title"`
	Url      string `json:"url" gorm:"column:url"`
	Rate     string `json:"rate" gorm:"column:rate"`
	PlayAble bool    `json:"playable" gorm:"column:play_able"`
	IsNew    bool    `json:"is_new" gorm:"column:is_new"`
	Cover    string `json:"cover" gorm:"column:cover"`
	CoverX   int64  `json:"cover_x" gorm:"column:cover_x"`
	CoverY   int64  `json:"cover_y" gorm:"column:cover_y"`
}

func (Movie) TableName() string {
	return "t_movies"
}
//tags
type Tags struct {
	Id int
	CatName  string `json:"id" gorm:"column:cate_name"`
	Sort  int `json:"sort" gorm:"column:sort"`
}
func (Tags) TableName() string {
	return "t_movie_tags"
}