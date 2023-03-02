/* package main

import "fmt"

func f0()  {			//	如何使用指针
	var a int = 20	//	声明实际变量
	var ip *int		//	声明指针变量

	ip = &a			//	指针变量的存储地址

	fmt.Printf("a 变量的地址是：%x\n",&a)
	
	fmt.Printf("ip 变量存储的指针地址是：%x\n", ip)

	fmt.Printf("ip 变量的值：%d\n", *ip)

}

func f1()  {
	var ptr *int
	fmt.Printf("ptr 的值为：%x\n", ptr)
}

func main() {
	f1()
	// f0()
	// var a int = 10
	// fmt.Printf("变量的地址：%x\n", &a)
} */