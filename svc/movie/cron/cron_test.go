package cron

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)
func TestUpdateMovie(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	//UpdateMovie(db,"日本")
}
func TestTags(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	SyncMocie(db)
	defer db.Close()
	select {

	}
}