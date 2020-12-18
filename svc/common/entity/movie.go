package entity
type Film struct {
	MovieId              int64    `json:"movie_id" db:"movie_id"`
	ActorName            []string `json:"actor_name" db:"actor_name"`
	Img                  string   `json:"img" db:"img"`
	Length               int64    `json:"length" db:"length"`
	IsSelectSeat         int64    `json:"is_select_seat" db:"is_select_seat"`
	FilmPrice            float32  `json:"film_price" db:"film_price"`
	FilmScreenwriter     string   `json:"film_screenwriter" db:"film_screenwriter"`
	CommentNum           int64    `json:"comment_num" db:"comment_num"`
	TitleCn              string   `json:"title_cn" db:"title_cn"`
	TitleEn              string   `json:"title_en" db:"title_en"`
	IsSupportInlineWatch int64    `json:"is_support_inline_watch" db:"is_support_inline_watch"`
	CreateAt             string   `json:"create_at" db:"create_at"`
	Type                 string   `json:"type" db:"type"`
	FilmDrama            string   `json:"film_drama" db:"film_drama"`
	CommonSpecial        string   `json:"common_special" db:"common_special"`
	UserAccessTimes      int64    `json:"user_access_times" db:"user_access_times"`
	FilmBoxoffice        float32  `json:"film_boxoffice" db:"film_boxoffice"`
	DirectorName         string   `json:"director_name" db:"director_name"`
	WantedCount          int64    `json:"wanted_count" db:"wanted_count"`
	UserCommentTimes     int64    `json:"user_comment_times" db:"user_comment_times"`
	CompanyIssued        string   `json:"company_issued" db:"company_issued"`
	Country              string   `json:"country" db:"country"`
	RatingFinal          float32  `json:"rating_final" db:"rating_final"`
	Is3D                 int64    `json:"is_3D" db:"is_3D"`
	IsDMAX               int64    `json:"is_DMAX" db:"is_DMAX"`
	IsFilter             int64    `json:"is_filter" db:"is_filter"`
	IsHot                int64    `json:"is_hot" db:"is_hot"`
	IsIMAX               int64    `json:"is_IMAX" db:"is_IMAX"`
	IsIMAX3D             int64    `json:"is_IMAX3D" db:"is_IMAX3D"`
	IsNew                int64    `json:"is_new" db:"is_new"`
	IsTicking            int64    `json:"is_ticking" db:"is_ticking"`
	RDay                 int64    `json:"r_day" db:"r_day"`
	RMonth               int64    `json:"r_month" db:"r_month"`
	RYear                int64    `json:"r_year" db:"r_year"`
	FilmDirector         string   `json:"film_director" db:"film_director"`
}
func (u Film) TableName() string {
	return "t_film"
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
//tags
type Tags struct {
	Id int
	CatName  string `json:"id" gorm:"column:cate_name"`
	Sort  int `json:"sort" gorm:"column:sort"`
}
func (Tags) TableName() string {
	return "t_movie_tags"
}