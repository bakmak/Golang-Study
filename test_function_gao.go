/* package main

import "fmt"

//	函数作为参数
func sayHello(name string)  {
	fmt.Printf("Hello,%s\n", name)
}

func f1(name string, f func(string) )  {
	f(name)
}

//	函数作为返回值
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) (func(int,int) int) {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil		
	}
}

func main() {
	f1("tow",sayHello)		//	函数作为参数
	add := cal("+")			//	函数作为返回值
	r := add(1,2)
	fmt.Printf("r: %v\n", r)
	fmt.Println("-----------------------")
	sub := cal("-")
	r = sub(100,50)
	fmt.Printf("r: %v\n", r)
} */