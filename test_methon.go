/* package main

import "fmt"

type Penson struct {
	name string
}

//	(per Penson) 接收者 receiver
func (per Penson) eat()  {
	fmt.Println(per.name + " eating...")
}

func (per Penson) sleep() {
	fmt.Println(per.name + " sleep...")
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string ) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	var per Penson
	per.name = "tom"
	per.eat()
	per.sleep()

	var customer = Customer{
		name : "tom",
	}
	i := customer.login("tom","123")
	fmt.Println(i)
}

 */