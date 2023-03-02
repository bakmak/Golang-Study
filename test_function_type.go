/* package main

import "fmt"

type fun func(int, int) int 		//	定义一个函数类型

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var f fun
	f = sum		//	相同函数类型，可以直接赋值
	s := f(1,2)
	fmt.Printf("s : %v\n", s)
	f = max
	m := f(3,4)
	fmt.Printf("m : %d\n", m)
} */