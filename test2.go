/* package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const ( 
	i = 1 << iota	// 0	1<<
	j 				// 1	2<<
	k 				// 1	3<<
	l				// 1	4<<
)

func f1() {
	fmt.Println("i = ",i)
	fmt.Println("j = ",j)
	fmt.Println("k = ",k)
	fmt.Println("l = ",l)
}

func f0()  {
	const(
		a = iota	// 0
		b			//1
		c			//2
		d = "ha"	//独立值，iota += 1
		e			//"ha"		iota += 1
		f = 100		//iota += 1
		g 			//100		iota += 1
		h = iota	//7,恢复计数
		i			//8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}



func main() {
	f1()
	f0()
	println(a , b , c)
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a , b , c = 1 , false , "str"		//	多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积 : %d", area)
	fmt.Println()
	fmt.Println(a,b,c)

} */