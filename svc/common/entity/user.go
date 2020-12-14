package entity

type User struct {
	UserId   int64  `json:"user_id" gorm:"primary_key"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password"  gorm:"column:password"`
	CreateAt string `json:"create_at" gorm:"column:create_at"`
	Email    string `json:"email"  gorm:"column:email"`
	Phone    string `json:"phone"  gorm:"column:phone"`
}

func (User) TableName() string {
	return "user"
}
