/* package main

import "fmt"

 func getNameAndge() (string,int) {
	return "阿达",20
} 

func main() {

	var name1 string
	var name string
	var age int
	var m bool 
	var(
		name string
		age intt
		m bool
	)  

	name = "tom"
	age = 20
	m = true
	
	var name string = "阿达"
	var site string = "www.duok360.com"
	var age int = 30
	
	// 类型推断 
	var name = "阿巴阿巴"
	var age = 20
	var site = "AVAB"


	// 批量初始化
	// var name , age , b = "tom" , 20 , true

	// 短变量声明
	 
	name := "tom"
	age := 20
	b := true 
	// _, name := getNameAndge()
	// fmt.Printf("name:%v\n",name)
	// fmt.Printf("b:%v\n",b)
	// fmt.Printf("m:%v\n",m)
	
	//**********************常量*********************************
	const PI float64 = 3.14
	const PI2 = 3.1415	//可以省略类型

	const (
		width = 100
		height = 200
	)

	const i, j = 1,2	//多重赋值
	const a,b,c = 1,2,"foo"

	fmt.Printf("PI: %v\n",PI)
	fmt.Printf("width: %v\n",width)
	fmt.Printf("height: %v\n",height)
	fmt.Printf("i: %v\n",i)
	fmt.Printf("j: %v\n",j)
	fmt.Printf("a: %v\n",a)
	fmt.Printf("b: %v\n",b)
	fmt.Printf("c: %v\n",c)
	
	const (
		A1 = iota
		B1 = iota
		C1 = iota	//类似i++
	)

	fmt.Printf("A1: %v\n",A1)
	fmt.Printf("B1: %v\n",B1)
	fmt.Printf("C1: %v\n",C1)

	const (
		A1 = iota
		_	//下划线就是跳过的意思
		_
		_
		_
		_
		_
		AN = iota
	)

	fmt.Printf("A1: %v\n",A1)
	fmt.Printf("AN: %v\n",AN)
	
	//iota中间插队，和下划线跳过的原理一样。
	const (
		A1 = iota
		A2 = 100
		A3 = iota

	)

 	fmt.Printf("A1: %v\n",A1)
	fmt.Printf("A2: %v\n",A2)
	fmt.Printf("A3: %v\n",A3) 

} */



