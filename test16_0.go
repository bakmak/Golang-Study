/* package main

import "fmt"

//	这里应该介绍一下 panic 与 recover,一个用于主动抛出错误，一个用于捕获panic抛出的错误。
//	panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误，panic 用于主动抛出错误，recover 用来捕获 panic 抛出的错误。

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	} ()
	defer func() {
		panic("three")
	} ()
	defer func() {
		panic("two")
	} ()
	panic("one")

} */