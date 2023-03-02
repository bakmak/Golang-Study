package main

import "fmt"

func f()  {
	var a = [5]int{1,2,3,4,5}
	for i, v := range a {
		fmt.Printf("i : %d, v : %v\n", i,v)
	}
}

func f1()  {
	var s = "多课网，go教程"
	for i, v := range s {
		fmt.Printf("i : %d, s : %c\n", i,v)
	}
}

func f2()  {
	var s = []int{1,2,3,4,5}
	for i, v := range s {
		fmt.Printf("i:%d, v: %v\n", i,v)
	}
}

func f3()  {
	m := make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"
	for k, v := range m {
		fmt.Printf("k : %v, v : %v\n", k,v)
	}
}

func main() {
	// f()
	// f1()
	// f2()
	f3()
}