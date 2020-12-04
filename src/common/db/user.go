package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mSystem/src/common/entity"
	"time"
)
// 通过ok
func InsertUser(userName string, password string, email string) error {

	today := time.Now().Format("2006-01-02")

	_, err := db.Exec("INSERT INTO `user`(`user_name`,`password`,`create_at`,`email`) VALUES(?,?,?,?)", userName, password, today, email)
	return err
}
// 通过ok
func SelectUserByEmail(email string) (*entity.User, error) {

	user := entity.User{}
	fmt.Print(db,email,"FFFFFFFFF")

	sqlStr := "SELECT * FROM user WHERE `email` = ?"
	e := db.QueryRow(sqlStr,email).Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateAt,&user.Email,&user.Phone)
	if e != nil {
		fmt.Printf("scan failed, err:%v\n", e)
		return nil, nil
	}
	return &user, e
}
//ok
func SelectUserByPasswordName(email string, password string) (*entity.User, error) {

	user := entity.User{}
	sql := "SELECT * FROM user WHERE `email` = ? AND `password` = ? LIMIT 1"
	e := db.QueryRow(sql,email,password).Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateAt,&user.Email,&user.Phone)
	if e != nil {
		fmt.Printf("scan failed, err:%v\n", e)
		return nil, nil
	}
	return &user, e
}
func SelectUserById(user_id float64,name string) (*entity.User, error) {

	user := entity.User{}
	sql := "SELECT * FROM user WHERE `user_id` = ? AND `user_name` = ? LIMIT 1"
	e := db.QueryRow(sql,int64(user_id),name).Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateAt,&user.Email,&user.Phone)
	if e != nil {
		fmt.Printf("scan failed, err:%v\n", e)
		return nil, nil
	}
	return &user, e
}
func UpdateUserNameProfile(userName string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `user_name` = ? WHERE user_id = ?", userName, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
//
func UpdateUserEmailProfile(email string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `email` = ? WHERE user_id = ?", email, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserPhoneProfile(phone string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE user_id = ?", phone, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
