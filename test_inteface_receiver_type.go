/* package main

import "fmt"

type Pet interface {		//	定义一个Pet接口
	eat()
}

type Dog struct {			//	定义一个Dog结构体
	name string
}

// func (dog Dog) eat()  {		//	实现Pet接口（接收者是值类）
// 	fmt.Printf("dog : %p\n",&dog)
// 	fmt.Println("dog eat...")
// 	dog.name = "黑黑"
// }

func (dog *Dog) eat()  {		//	将Pet接口改为指针接收者
	fmt.Printf("dog : %p\n",dog)
	fmt.Println("dog eat...")
	dog.name = "黑黑"
}

func main() {
	//	接口值类型接收者
	// dog := Dog{name : "花花"}
	// fmt.Printf("dog : %p\n", &dog)
	// dog.eat()
	// fmt.Printf("dog : %v\n", dog)

	//	接口指针类型接收者
	dog := Dog{name : "花花"}
	fmt.Printf("dog : %p\n", &dog)
	dog.eat()
	fmt.Printf("dog : %v\n", dog)
} */