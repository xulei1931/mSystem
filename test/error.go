package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	Name string
	time.Time
	Err error
}

func (m MyError) Error() string {
	return fmt.Sprintf("%v %v %v", m.Name, m.Time, m.Err)
}
func (m MyError) Unwrap() error {
	return m.Err
}

func TestErr() error {
	return MyError{
		Name: "AAAA",
		Time: time.Now(),
	}
}


func main() {
	e := errors.New("原始错误e")

	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w,e)) // 是否w包含了e的原始错误
	fmt.Println(errors.As(w,e))
	fmt.Println(e)
}
