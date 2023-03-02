/* package main

import "fmt"

//	接口实例
// type USB interface {		//	USB接口
// 	read()
// 	write()
// }

// type Computer struct {	//	Computer接口
// }

// type Mobile struct {		//	Mobile接口
// }
//	Computer实现USB接口
// func (c Computer) read()  {			
// 	fmt.Println("computer read...")
// }

// func (c Computer) write()  {			
// 	fmt.Println("computer write...")
// }
//	Mobile实现USB接口
// func (c Mobile) read()  {
// 	fmt.Println("mobile read...")
// }

// func (c Mobile) write()  {
// 	fmt.Println("mobile write...")
// }

//	实现接口必须实现接口中的所有方法
type OpenClose interface {
	open()
	close()
}

type Door struct {
}

func (d Door) open()  {
	fmt.Println("open dor...")
}

func main() {
	// c := Computer{}
	// m := Mobile{}
	// c.read()
	// c.write()
	// m.read()
	// m.write()

	var oc OpenClose
	oc = Door{}	// 这里编译错误，提示只实现了一个接口
} */