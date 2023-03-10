/* package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) (string,error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s,nil
	}
}

//	MyError是一个包含时间和消息的错误实现。
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file systme has gone away",
	}
}

func main() {
	// s, err := check("")
	// if err != nil {
	// 	fmt.Printf("err.Error(): %v\n", err.Error())
	// } else {
	// 	fmt.Printf("s: %v\n", s)
	// }
	err := oops()
	if err != nil {
		fmt.Println(err)
	}

} */