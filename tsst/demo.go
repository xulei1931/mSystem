package main

import (
	"fmt"
	"strconv"
	"time"
)
/**
获取本周周一的日期
*/

func GetFirstDateOfWeek() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

}

/**
获取上周的周一日期
*/
//func GetLastWeekFirstDate() (weekMonday string) {
//	thisWeekMonday := GetFirstDateOfWeek()
//	fmt.Println(thisWeekMonday)
//	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
//	lastWeekMonday := TimeMonday.AddDate(0, 0, 7)
//	weekMonday = lastWeekMonday.Format("2006-01-02")
//	return
//}
func main()  {
	fmt.Println(GetFirstDateOfWeek())
	//fmt.Println(time.Now())
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 9.824), 64)
	fmt.Println(value) //9.82
}