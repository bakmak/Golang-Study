/* package main

import "fmt"

func f()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Printf("i : %v\n", i)
		}
	}
}

func f1()  {
	for i := 0; i < 5; i++ {
		MY_LABEL:
			for j := 0; j < 5; j++ {
				if i == 2 && j == 2 {
					continue MY_LABEL
				}
				fmt.Printf("i = %d, j = %d\n",i,j)
			}
	}
}

func f2(){
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("i: %v\n", i)
		}else {
			continue
		}
	}
}

func main() {
	// f()
	// f1()
	f2()
} */