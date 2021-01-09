package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	a1:=[2]string{"飞雪无情","张三"}

	s1:=a1[0:1]

	s2:=a1[:]

	//打印出s1和s2的Data值，是一样的

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

}
