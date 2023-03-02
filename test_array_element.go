/* package main

import "fmt"

func f()  {
	var a [5]int
	a[0] = 100
	a[3] = 20

	fmt.Printf("a : %v\n", a)
	fmt.Println("--------------------------------------")
	a[0] = 40
	a[3] = 60

	fmt.Printf("a : %v\n", a)
}

func f1()  {
	a := [...]int{1,2,3,4,5,6}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] : %v\n", i,a[i])
	}
}

func f2()  {
	a := [...]int{1,2,3,4,5,6}
	for i, v := range a {
		fmt.Printf("a[%d]: %v\n", i,v)
	}
}

func main() {
	// f()
	// f1()
	f2()
} */