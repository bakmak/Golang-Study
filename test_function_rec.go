/* package main

import "fmt"

func f(n int) int {									//	斐波那契数列
	if n == 1 || n ==2 {	//	退出点判断
		return 1
	}
	return f(n-1) + f(n-2)		//	递归表达式
}

func a(n int) int {									//	阶乘
	//	5的阶乘就是54321 5x4x3x2x1	下面是for循环方法
	// var s int = 1
	// for i := 1; i <= n; i++ {
	// 	s *= i
	// }
	// return s
	
	if n == 1 {		//	返回条件
		return 1	
	} else {		//	自己调用自己
		return n * a(n-1)
	}
}

func main() {
	n := 5
	r := a(n)
	fmt.Printf("r : %v\n", r)

	r = f(5)
	fmt.Printf("r : %v\n", r)
} */