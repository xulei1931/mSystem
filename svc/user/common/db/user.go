package db

import (
	"encoding/json"
	"fmt"
	"user/common/entity"
	"time"
)
func init(){
	 //if err := recover();err!=nil{
	 //	fmt.Println(err)
	 //}
}
// 通过ok
func InsertUser(userName string, password string, email string) error {

	today := time.Now().Format("2006-01-02")
	u := new(entity.User)
	u.Email=email
	u.UserName=userName
	u.Password=password
	u.CreateAt=today
	fmt.Println(u)

	err := db.Create(&u)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 通过ok
func SelectUserByEmail(email string) bool {

	user := entity.User{}
	er := db.Where("email=?", email).Find(&user)
	if er.Error != nil{
		return false
	}
	return true
}

//ok
func SelectUserByPasswordName(email string, password string) (*entity.User, error) {

	user := new(entity.User)
	er := db.Where("email=? and password=? ", email,password).Find(&user)

	if er.Error !=nil{ // 差不到也走这里

		return user, er.Error

	}
	return user, nil
}
func SelectUserById(user_id float64, name string) (*entity.User, error) {
	user := entity.User{}
	user_key := fmt.Sprintf("user_id:", user_id)
	// redis
	redis_values, err := redisClient.Get(user_key).Result()
	if err == nil {
		b := []byte(redis_values)
		if json.Unmarshal(b, &user) == nil {
			fmt.Printf("get from redis value :%v\n", &user)
			return &user, nil
		}
	}
	db.Where("user_id=? and user_name=? ", user_id,name).Find(&user)
	bytes1, err := json.Marshal(&user)
	if err == nil {
		// 返回的是字节数组 []byte
		redisClient.Set(user_key, string(bytes1), 600) // 10分钟
	}
	return &user, nil
}
func UpdateUserNameProfile(userName string, userId int64) error {

	return nil
}

//
func UpdateUserEmailProfile(email string, userId int64) error {

	return nil
}

func UpdateUserPhoneProfile(phone string, userId int64) error {

	return nil
}
