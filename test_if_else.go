/* package main

import "fmt"

func f1() {
	a := 1
	b := 2
	if a > b {
		fmt.Printf("\"a>b\": %v\n", a > b)
	} else {
		fmt.Printf("\"a<b\": %v\n", a < b)
	}
}

func f2() {
	var s int
	fmt.Println("输入一个数字：")
	fmt.Scan(&s)

	if s%2 == 0 {
		fmt.Println("s 是偶数")
	} else {
		fmt.Println("s 是奇数")
	}
	fmt.Printf("s 的值是：%v\n", s)
}

func f3() {
	age := 20
	if age >= 18 {
		fmt.Println("你是成年人")
	} else {
		fmt.Println("你还未成年")
	}
}

func f4()  {
	if age := 20; age >= 18 {
		fmt.Println("你可以看东京热了")
	} else {
		fmt.Println("你还不能看东京热")
	}
}

func main() {
	f1()
	f2()
	f3()
	f4()
	var name string
	var age string
	var email string
	fmt.Println("请输入name，age，email，用空格进行分割")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name : %v\n", name)
	fmt.Printf("age : %v\n", age)
	fmt.Printf("email : %v\n", email)

}
 */