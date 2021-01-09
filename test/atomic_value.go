package main

import (
	"fmt"
	"sync/atomic"
)

//Mutex由操作系统实现，而atomic包中的原子操作则由底层硬件直接提供支持。
//在 CPU 实现的指令集里，有一些指令被封装进了atomic包，这些指令在执行的过程中是不允许中断（interrupt）的，
//因此原子操作可以在lock-free的情况下保证并发安全，并且它的性能也能做到随 CPU 个数的增多而线性扩展
func main(){
	var config atomic.Value
	data := make(map[int]string)
	data[1]="aaa"
	config.Store(data)
	fmt.Println(config.Load().(map[int]string))
	//
	//bytes := []byte{104, 101, 108, 108, 111}
	//fmt.Println(&bytes)
	//p := unsafe.Pointer(&bytes)
	//strp:= (*string)(p)
	//fmt.Println(p,*strp)
}