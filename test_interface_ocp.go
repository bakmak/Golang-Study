/* package main

import "fmt"

type Pet interface {			//	定义一个宠物接口Pet
	eat()
	sleep()
}

type Dog struct {				//	定义一个Dog结构体
	name string
	age int
}

//	实现这个接口就可以当Pet使用
func (dog Dog) eat()  {			//	Dog实现Pet接口方法
	fmt.Println("dog eat...")
}

func (dog Dog) sleep()  {		//	Dog实现Pet接口方法
	fmt.Println("dog sleep...")
}

type Cat struct {				//	定义一个Cat结构体
	name string
	age int 
}

//	实现这个接口就可以当Pet使用
func (cat Cat) eat()  {			//	Cat实现Pet接口方法
	fmt.Println("cat eat...")
}

func (cat Cat) sleep()  {		//	Cat实现Pet接口方法
	fmt.Println("cat sleep...")
}

type Person struct {			//	定义一个Person结构体
	name string
}

func (per Person) care(pet Pet)  {			//	为Person添加一个养宠物方法
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	per := Person{}

	per.care(dog)
	per.care(cat)
} */