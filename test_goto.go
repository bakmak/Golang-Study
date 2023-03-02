/* package main

import "fmt"

func f()  {
	a := 1
	if a == 0 {
		goto LABEL1
	} else {
		fmt.Println("other")
	}
	LABEL1:
		fmt.Println("next...")
}
func f1()  {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 10; k++ {
				if i == 2 && j == 2 {
					goto LABEL1
				}
				fmt.Printf("i : %v \t j : %v \t k : %v\n", i,j,k)	
			}
		}
	}
	LABEL1:
		fmt.Println("label1")
}


func main() {
	// f()
	f1()
} */