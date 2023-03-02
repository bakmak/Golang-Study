/* package main

import "fmt"

func f1(a int, b int) int {			//	传参
	if a > b {
		return a
	} else {
		return b
	}
}

func f2(a int)  {					//	按值传递
	a = 200
	fmt.Printf("a1 : %v\n", a)
}

func f3(a []int)  {
	a [0] = 1000
}

func f4(args ...int) {
	for _, v := range args {
		fmt.Printf("v : %v\n", v)
	}
}

func f5(name string, age int, ages ...int)  {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range ages {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// r := f1(1,2)
	// fmt.Printf("r : %v\n", r)
	// a := 100
	// f2(a)
	// fmt.Printf("a : %v\n", a)	

	// a1 := []int{1,2}
	// fmt.Printf("a1 : %v\n ", a1)
	// f3(a1)
	// fmt.Printf("a1 : %v\n ", a1)

	f4(1,2,3)
	fmt.Println("-------------------------")
	f4(1,2,3,4,5,6)
	fmt.Println("-------------------------")
	f5("tow",20,5,6,7,8)
} */