/* package main

import "fmt"

func f0()  {	//	算术运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n",c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n",c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n",c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n",c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n",c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n",a)
	a = 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

func f1() {			//		关系运算符
	var a int = 21
	var b int = 10
	if ( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}

	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}

	a = 5
	b = 20

	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n")
	} 
	if ( a >= b ) {
		fmt.Printf("第五行 - a 大于等于 b\n")
	}
}

func f2()  {		//	逻辑运算符
	var a bool = true
	var b bool = false
	if ( a && b ) {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if ( a || b ) {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	a = false
	b = true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if ( !(a && b) ) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

func f3()  {		//		位运算符
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

func f4()  {	//		赋值运算符
	var a int = 21 
	var c int

	c = a
	fmt.Printf("第1行 - = 运算符实例，c 值为 = %d\n",c)

	c += a
	fmt.Printf("第2行 - += 运算符实例，c 值为 = %d\n",c)

	c -= a
	fmt.Printf("第3行 - -= 运算符实例，c 值为 = %d\n",c)

	c *= a
	fmt.Printf("第4行 - *= 运算符实例，c 值为 = %d\n",c)

	c /= a
	fmt.Printf("第5行 - /= 运算符实例，c 值为 = %d\n",c)

	c = 200

	c <<= 2
	fmt.Printf("第6行 - <<= 运算符实例，c 值为 = %d\n",c)

	c >>= 2
	fmt.Printf("第7行 - >>= 运算符实例，c 值为 = %d\n",c)

	c &= 2
	fmt.Printf("第8行 - &= 运算符实例，c 值为 = %d\n",c)

	c ^= 2
	fmt.Printf("第9行 - ^= 运算符实例，c 值为 = %d\n",c)

	c |= 2
	fmt.Printf("第10行 - |= 运算符实例，c 值为 = %d\n",c)
}

func f5()  {	//	其他运算符
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第1行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第2行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第3行 - c 变量类型为 = %T\n", c)

	ptr = &a
	fmt.Printf("a 的值为 %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)	
}

func f6()  {		//		运算符优先级
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d
	fmt.Printf("((a + b) * c) / d 的值为 : %d\n", e)

	e = (a + b) * (c / d)
	fmt.Printf("(a + b) * (c / d) 的值为 : %d\n", e)

	e = a + (b * c) / d
	fmt.Printf("a + (b * c) / d 的值为 : %d\n", e)

}

func main() {
	// f0()
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
} */

package main

import "fmt"

func main() {
	/* 使用指针变量与不使用的区别 */ 
	var a int = 4
	var prt int
	prt = a 
	fmt.Println(prt)	//	4
	a = 15
	fmt.Println(prt)	//	4

	var b int = 5
	var prt1 *int
	prt1 = &b
	fmt.Println(*prt1)	//	5
	b = 15
	fmt.Println(*prt1)	//	15

	/* 指针变量*和地址值&的区别 */
	// var a int = 4
	// var ptr *int
	// ptr = &a
	// fmt.Println("a 的值为",a)
	// fmt.Println("*ptr为",*ptr)
	// fmt.Println("ptr为",ptr)
}




