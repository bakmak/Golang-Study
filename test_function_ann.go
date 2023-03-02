/* package main

import "fmt"
//	匿名函数有什么作用？就是在函数内部做一些运算
func test()  {
	name := "tow"
	age := "20"
	f1 := func () string {
		return name + age
	}
	msg := f1()
	fmt.Printf("msg : %v\n", msg)
}

func main() {
	//	匿名函数就没有名称
	// max := func (a int, b int) int {		//	匿名函数实例
	// 	if a > b {
	// 		return a 
	// 	} else {
	// 		return b
	// 	}
	// }

	// i := max(1,2)
	// fmt.Printf("i : %v\n", i)

	func (a int, b int)  {			//	自己执行
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max)
	}(1,2)

	test()
} */