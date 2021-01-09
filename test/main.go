package main

import (
	"fmt"
	"time"
	"unsafe"
)
const TimeLayout = "2006-01-02 15:04:05"
type person struct {
	Name string
	Age int
}

func main(){
	p :=new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName:=(*string)(unsafe.Pointer(p))
	*pName="飞雪无情"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	fmt.Println(uintptr(unsafe.Pointer(p)),unsafe.Offsetof(p.Age))
	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)

	fmt.Println(unsafe.Sizeof(true))//1
	fmt.Println(unsafe.Sizeof(int8(0)))//1
	fmt.Println(unsafe.Sizeof(int16(10))) //2
	fmt.Println(unsafe.Sizeof(int32(10000000))) //4
	fmt.Println(unsafe.Sizeof(int64(10000000000000))) //8
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))//8
	fmt.Println(unsafe.Sizeof(string("飞雪无情"))) //16
	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情","张三"})) //24
}
func test1()  {
	localUTC, _ := time.LoadLocation("Local")
	//loc, _ := time.LoadLocation("PRC")
	//t, _ := time.ParseInLocation("2006-01-02", "2018-05-31", loc)
	//fmt.Println(t)
	tTime, _ := time.ParseInLocation(TimeLayout, "2019-11-15 23:12:25", localUTC)
    fmt.Println(tTime)
}
