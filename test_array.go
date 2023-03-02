/* package main

import "fmt"

func f()  {
	var a [3]int
	var s [2]string
	var b [2]bool
	fmt.Printf("a : %v \t type: %T\n", a,a)
	fmt.Printf("s : %v \t type: %T\n", s,s)
	fmt.Printf("b : %v \t type: %T\n", b,b)
}

func f1()  {
	var a = [3]int{1,2,3}
	var s = [2]string{"tow","kite"}
	var b = [2]bool{true,false}
	a1 := [2]int{1,2}
	fmt.Printf("a : %v\n", a)
	fmt.Printf("s : %v\n", s)
	fmt.Printf("b : %v\n", b)
	fmt.Printf("a1 : %v\n", a1)
}

func f2()  {
	var a = [...]int{1,2}
	var s = [...]string{"tow","kite"}
	var b = [...]bool{true,false}
	a1 := [...]int{1,2}
	fmt.Printf("a : %v\n", a)
	fmt.Printf("s : %v\n", s)
	fmt.Printf("b : %v\n", b)
	fmt.Printf("a1 : %v\n", a1)
}

func f3()  {
	var a = [...]int{5:1,9:2}
	var s = [...]string{2:"tow",5:"kite"}
	var b = [...]bool{1:true,3:false}
	a1 := [...]int{5:1,9:5}
	fmt.Printf("a : %v\nlen(a):%v\n", a,len(a))
	fmt.Printf("s : %v\nlen(s):%v\n", s,len(s))
	fmt.Printf("b : %v\nlen(b):%v\n", b,len(b))
	fmt.Printf("a1 : %v\nlen(a1):%v\n", a1,len(a1))
}

func main() {
	f()
	fmt.Println("--------------------------------------------------------------------------\n")
	f1()
	fmt.Println("--------------------------------------------------------------------------\n")
	f2()
	fmt.Println("--------------------------------------------------------------------------\n")
	f3()
	
} */