/* package main

import "fmt"

type Animal struct {
	name string
	age int
}

func (a Animal) eat()  {
	fmt.Println("eat...")
}

func (a Animal) sleep()  {
	fmt.Println("sleep")
}

type Dog struct {
	Animal					//	可以理解为继承关系
}

type Cat struct {
	Animal					//	可以理解为继承关系
}

func main() {
	dog := Dog{
		Animal{		//	匿名方式定义声明这个Animal
			name : "dog",
			age : 2,
		},
	}

	cat := Cat{
		Animal{
			name : "cat", age : 3},		// 写在一行 3 后面不用加 , 逗号
	}

	//	dog和cat可以同时调用Anilmal方法
	dog.eat()
	dog.sleep()
	//	使用的匿名方式调用这个Anilmal
	cat.eat()
	cat.sleep()

} */