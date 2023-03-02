/* package main

import "fmt"

func testAppend() {
	s1 := []int{1, 2, 3}
	i := append(s1, 4)
	fmt.Printf("i: %v\n", i)

	s2 := []int{7, 8, 9}
	i2 := append(s1, s2...)
	fmt.Printf("i2: %v\n", i2)
}

func testLen()  {
	u := "hello world"
	i := len(u)
	fmt.Printf("i: %v\n", i)

	i2 := []int{1, 2, 3}
	i3 := len(i2)
	fmt.Printf("i3: %v\n", i3)
}

func testPrint()  {
	name := "tom"
	age := 20
	print(name, " ",age , "\n")
	fmt.Println("------------------------------")
	println(name, " ", age)
}

func testPanic()  {
	defer fmt.Println("panic 异常后执行...")
	panic("panic 错误...")
	fmt.Println("end...")
}

func testNew()  {
	b := new(bool)
	fmt.Println(*b)
	i := new(int)
	fmt.Println(*i)
	s := new(string)
	fmt.Println(*s)
}


func testMakeVsNew()  {
	var p *[]int = new([]int)
	v := make([]int, 10)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("v: %v\n", v)
}

func main() {
	// testAppend()
	// testLen()	
	// testPrint()
	// testPanic()
	// testNew()
	testMakeVsNew()
} */