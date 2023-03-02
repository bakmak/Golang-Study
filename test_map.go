/* package main

import "fmt"

func test0()  {		//使用内建函数make进行定义map变量
	m1 := make(map[string]string)
	m1["name"] = "tow"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"
	fmt.Printf("m1 : %v", m1)
}

func test1()  {		//声明变量，map默认是nil
	var m2 map[string]string
	fmt.Printf("m2 : %v\n", m2)
	m2 = map[string]string{
		"name" : "kite" ,
		"age" : "20",
		"email" : "tom@gmail.com",
	}
	fmt.Printf("m2 : %v\n", m2)
}

func test2()  {		//		访问map
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	name := m1["name"]
	age := m1["age"]
	email := m1["email"]
	fmt.Printf("name : %v\n", name)
	fmt.Printf("age : %v\n", age)
	fmt.Printf("email : %v\n", email)
}

func test3()  {		//	判断某个键是否存在
	m1 := make(map[string]string)
	m1["name"] = "tow"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	v, ok := m1["tow"]
	if ok {
		fmt.Println("键存在")
		fmt.Println(v)
	} else {
		fmt.Println("键不存在")
	}

}

func test4()  {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

func main() {
	// test0()
	// test1()
	// test2()
	// test3()
	test4()
} */