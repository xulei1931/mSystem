package main

import (
	"fmt"
)

func main() {
	//Go(go_test)
	go_test()

	//select {
	//}
}
func go_test() {
	fmt.Println("go......")
	//panic("22")
}
func Go(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go f()
}
