package main

import (
	"fmt"
	"time"
)

func main()  {
	 now := time.Now()
	 fmt.Println(now.String())
	//t, _ := time.Parse("2006-01-02 15:04:05", now.String())
	t, _ := time.ParseInLocation("2006-01-02 15:00:00", now.String(),time.Local)
	fmt.Println(time.Now().Sub(t).Hours())

	fmt.Println(now.Format("2006-01-02 15:00:00"))


	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t1)
	test()
}
func test(){
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(t.Round(1 * time.Hour))

	// 整分（向下取整）
	fmt.Println(t.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(t.Round(1 * time.Minute))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t2)
}