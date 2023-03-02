/* //	闭包其实并不复杂，只要牢记 “ 闭包 = 函数 + 引用环境 ”
package main

import (
	"fmt"
	"strings"
)

func add() func(int) int {			//	返回一个函数
	var x int
	return func (y int) int {
		fmt.Printf("x : %v\n",x)
		x += y
		fmt.Printf("x + %v = %v\n",y, x)
		return x	
	}
}

func add0(x int) func(int) int {
	fmt.Printf("x : %v\n",x)
	return func(y int) int {
		fmt.Printf("y : %v\n", y)
		x += y
		fmt.Printf("x + y : %v\n", x)
		return x
	}
}

func makeSuffixFunc(suffix string) func (string) string {
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}


func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// var f = add()
	// fmt.Printf("f : %v\n",f(10))
	// fmt.Println("------------------------")
	// fmt.Printf("f : %v\n",f(20))
	// fmt.Println("------------------------")
	// fmt.Printf("f : %v\n",f(30))
	// fmt.Println("------------------------")
	// f1 := add()
	// fmt.Printf("f : %v\n", f1(40))
	// fmt.Println("------------------------")
	// fmt.Printf("f : %v\n", f1(50))
	// fmt.Println("------------------------")

	// var f = add0(0)
	// fmt.Println("--------------")
	// fmt.Println(f(10))
	// fmt.Println("--------------")
	// fmt.Println(f(20))
	// fmt.Println("--------------")
	// fmt.Println(f(30))
	// fmt.Println("---------------------------")
	// f1 := add0(20)
	// fmt.Println("--------------")
	// fmt.Println(f1(40))
	// fmt.Println("--------------")
	// fmt.Println(f1(50))

	// jpgFunc := makeSuffixFunc(".jpg")
	// txtFunc := makeSuffixFunc(".txt")
	// fmt.Println(jpgFunc("test"))
	// fmt.Println(txtFunc("test"))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
} */