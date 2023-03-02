/* package main

import "fmt"

type Penson struct {
	name string
}

// func showPerson(per Penson)  {			//	值传递
// 	fmt.Printf("per: %p\n", &per)
// 	per.name = "kite"
// 	fmt.Printf("per: %v\n", per)
// }

// func showPerson2(per *Penson)  {		//	地址传递
// 	fmt.Printf("per: %p\n", per)
// 	per.name = "kite"
// 	fmt.Printf("per: %v\n", per)
// }

func (per Penson) showPerson()  {
	fmt.Printf("per: %p\n", &per)
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func (per *Penson) showPerson2()  {
	fmt.Printf("per: %p\n", per)
	per.name = "kite"
	fmt.Printf("per: %v\n", *per)
}

func main() {
	p1 := Penson{name : "tom"}		
	fmt.Printf("p1: %p\n", &p1)
	p1.showPerson()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("-------------------")
	p2 := &Penson{name : "tom"}		
	fmt.Printf("p2: %p\n", p2)
	p2.showPerson2()
	fmt.Printf("p2: %v\n", *p2)
	// p1 := Penson{name : "tom"}		//	值类型
	// fmt.Printf("p1: %p\n", &p1)
	// showPerson(p1)
	// fmt.Printf("p1: %v\n", p1)
	// fmt.Println("-------------------")
	// p2 := &Penson{name : "tom"}		//	指针类型
	// fmt.Printf("p2: %p\n", p2)
	// showPerson2(p2)
	// fmt.Printf("p2: %v\n", p2)

} */