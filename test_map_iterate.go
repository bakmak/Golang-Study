package main

import "fmt"

func test0()  {		//		遍历key
	m := make(map[string]string)
	m["name"] = "tow"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"

	for key := range m {
		fmt.Println(key)
	}
}

func test1()  {		//		遍历key和value
	m := make(map[string]string)
	m["name"] = "tow"
	m["age"] = "20"
	m["email"] = "tow@gmail.com"

	for key, v := range m {
		fmt.Println(key + ":" + v)
	}
}

func main() {
	// test0()
	test1()
}