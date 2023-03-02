/* package main

import "fmt"

var x, y int
var ( //这种因式分解关键字的写法一般用于声明全局变量
	a int
	b int
)

var c, d int = 1, 2
var e, f = 123, "hello"

//	这种不带声明格式的只能在函数体中出现
//	g , h := 123 , "hello"

func f4() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}

func f0() {
	var a = "Runoob"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)

}

func f1() { //	1.指定变量类型，如果没有初始化，则变量默认为零
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func f2() { //	2.根据值自行判断变量类型
	var d = true
	fmt.Println(d)
}

func f3() { //	3.如果变量已经使用var声明过了，再使用 := 声明变量，就会产生错误。
	f := "Runoob" // == var f string = "Runoob"
	fmt.Println(f)
}

func numbers() (int, int, string) { //	一个可以返回多个值的函数
	a, b, c := 1, 2, "str"
	return a, b, c
}

func main() {
	_, number, strs := numbers()
	fmt.Println(number, strs)
	// f4()
	// f3()
	// f2()
	// f1()
	// f0()
	// var a string = "Runoob"
	// fmt.Println(a)

	// var b , c int = 1 , 2
	// fmt.Println(b,c)
}
 */