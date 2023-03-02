/* package main

import "fmt"

func f1() {			//	没有返回值
	fmt.Printf("我没有返回值，只是进行一些计算\n")
}

func sum(a int, b int) (ret int) {			//	有返回值
	ret = a + b
	return ret
}

func f2() (name string, age int) {			//	多个返回值，且在return中指定返回的内容
	name = "阿达"
	age = 24
	return name,age
}

func f3() (name string, age int) {			//	多个返回值，返回值名称没有被使用
	name = "巴卡"
	age = 14
	return			//	等价于 return name,age
}

func f4() (name string, age int) {			//	return覆盖命名返回值，返回值名称没有被使用
	n := "老郭"
	a := 30
	return n,a
}

func f5() (string,int) {
	n := "阿😓"
	a := 20
	return n,a
}

func main() {
	f1()
	fmt.Printf("有返回值：%v\n", sum(1,2))
	name,age := f2()
	fmt.Printf("多个返回值，且在return中指定返回的内容：%v\t%v\n", name,age)
	name1 , age1 := f3()
	fmt.Printf("多个返回值，返回值名称没有被使用：%v \t %v\n", name1,age1)
	n,a := f4()
	fmt.Printf("return覆盖命名返回值，返回值名称没有被使用；%v \t %v \n", n,a)
	n1,a1 := f5()
	fmt.Printf("形参没有名称，返回值名称没有被使用：%v \t %v\n", n1,a1)
} */