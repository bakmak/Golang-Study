/* 
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"go/format"
	//"strings"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test0() {			//	结构体转换为json
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("string(b): %v\n", string(b))
}


func test1()  {			//	json转换为结构体
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var m Person
	json.Unmarshal(b, &m)
	fmt.Printf("m: %v\n", m)
}

func test2()  {			//	解析嵌套类型
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["tom", "kite"]}`)
	// var f interface{}
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f)

	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

func test3()  {			//	解析嵌套引用类型
	type Person	struct{
		Name 	string
		Age		int
		Email	string
		Parent 	[]string
	}

	p := Person{
		Name: "tom",
		Age: 20,
		Email: "tom@email",
		Parent: []string{"big tom", "big kite"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test4()  {
	f4, _ := os.Open("a.json")
	d := json.NewDecoder(f4)
	e := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := d.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("v: %v\n", v)
		if err := e.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

func test5()  {
	f4, _ := os.Open("a.json")
	defer f4.Close()
	d := json.NewDecoder(f4)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Printf("v: %v\n", v)
}

func test6()  {
	type Person struct {
		Name	string
		Age		int
		Email	string
		Parent	[]string
	}

	p := Person{
		Name:	"AAA",
		Age: 20,
		Email: "AAA@Email.com",
		Parent: []string{"big AAA", "big BBB"},
	}

	f4, _ := os.OpenFile("a.json", os.O_WRONLY, 0777)
	defer	f4.Close()
	e := json.NewEncoder(f4)
	e.Encode(p)
}

func main() {
	// test0()
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()
} */