/* package main

import "fmt"

type Reader interface {		//	interface本质就是一个指针
	ReadBook()
}

type Writer interface {
	WirteBook()
}

type Book struct {
}

func (t *Book) WirteBook()  {
	fmt.Println("write a book")
}

func (t *Book) ReadBook()  {
	fmt.Println("read a book")
}

func main() {
	b := &Book{}	//	b : pair<type:Book, value:book{}地址>

	var r Reader	//	r : pair<type, value>
					//	r : pair<type:Book, value:book{}地址>	pair是不变的
	r = b			//	interface r 本质就是一个指针，所以b也要是指针类型	
	r.ReadBook()	

	var w Writer	// w : pair<type:Book, value:book{}地址>    pair是不变的
	w = r.(Writer)	// w和r的type一致，所以断言成功（断言就是强转？）

	w.WirteBook()
} */