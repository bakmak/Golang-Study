/* package main

import "fmt"

var g int = 20

var a int = 20	

var Total_sum int = 0

func sum(a, b int) int {	//	函数定义-两数相加
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}

func f1()  {
	var a, b int
	a = 10
	b = 20
	g = a + b
	fmt.Printf("结果：a = %d， b = %d and g = %d\n", a, b, g)
}

func f0()  {	//	局部变量
	var a, b, c int		//	声明局部变量
	a = 10				//	初始化参数
	b = 20
	c = a + b
	fmt.Printf("结果：a = %d, b = %d and c = %d\n", a, b, c)
}

func f2 () {
	var a int = 0
	fmt.Println("for start")
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("for end")

	fmt.Println(a)
}

func f3()  {
	var a int = 0
	fmt.Println("for start")
	for a = 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("for end")

	fmt.Println(a)
}

func Sum_test(a int, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	Total_sum += (a + b)
	fmt.Printf("Total_sum：%d\n", Total_sum)
	return a + b
}

func main() {
	// f0()
	// f1()
	// var g int = 10
	// fmt.Printf("结果：g = %d\n", g)		//	全局变量与局部变量名称可以相同，但函数内的局部变量会被优先考虑
	
	// var a int = 10
	// var b int = 20
	// var c int = 0
	// fmt.Printf("main()函数中 a = %d\n", a)
	// c = sum( a, b)							// 使用了局部变量a
	// fmt.Printf("main()函数中 c = %d\n", c)
	// f2()
	// f3()

	var sum int 
	sum = Sum_test(2,3)
	fmt.Printf("sum：%d；Total_sum：%d\n", sum, Total_sum)

}

 */