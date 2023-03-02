/* package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func changeBook(book *Books)  {
	book.title = "book_change"
}

func printBook( book *Books )  {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func f1()  {
	var Book1 Books		//	声明Book1为Books类型
	var Book2 Books		//	声明Book2为Books类型

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Pyramid 语言教程"
	Book2.book_id = 6495700

	// printBook(&Book1)
	// fmt.Printf("Book1 title : %s\n", Book1.title)
	// fmt.Printf("Book1 author : %s\n", Book1.author)
	// fmt.Printf("Book1 subjectr : %s\n", Book1.subject)
	// fmt.Printf("Book1 book_id : %d\n", Book1.book_id)

	// printBook(&Book2)
	// fmt.Printf("Book2 title : %s\n", Book2.title)
	// fmt.Printf("Book2 author : %s\n", Book2.author)
	// fmt.Printf("Book2 subjectr : %s\n", Book2.subject)
	// fmt.Printf("Book2 book_id : %d\n", Book2.book_id)

	// changeBook(&Book1)
	// fmt.Println(Book1)

	var b *Books

	b = &Book1

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)
	fmt.Println(Book1)
}

func f0()  {
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})		//	创建一个新的结构体
	fmt.Println(Books{title:"Go 语言", author:"www.runoob.com", subject:"Go 语言教程", book_id:6495407})	//也可以使用 key => value 格式
	fmt.Println(Books{title:"Go 语言", author:"www.runoob.com"})
}

func main() {
	// f0()
	f1()
} */