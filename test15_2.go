/* package main

import "fmt"

type Phone interface {		//	定义接口
	call()
	call2()
}

type Phone1 struct {			//	一直都搞不懂这是干啥的，原来是用来定义结构体内的数据类型的
	id				int
	name			string
	category_id		int
	category_name	string
}

func (test Phone1) call()  {		//	第一个类的第一个回调函数
	fmt.Println("这是第一个类的第一个接口回调函数 结构体数据：",Phone1{id: 1, name: "浅笑"})
}

func (tets Phone1) call2()  {	//	第一个类的第二个回调函数
	fmt.Println("这是一个类的第二个接口回调函数call2", Phone1{id: 1, name: "浅笑", category_id: 4, category_name: "分类名称"})
}

type Phone2 struct {			//	第二个结构体的数据类型
	member_id			int
	member_balance		float32
	member_sex			bool
	member_nickname		string
}

func (test2 Phone2) call()  {	//	第二个类的第一个回调函数
	fmt.Println("这是第二个类的第一个接口回调函数call", Phone2{member_id: 22, member_balance: 15.23, member_sex: false, member_nickname: "浅笑18"})
}

func (test2 Phone2) call2()  {	//	第二个类的第二个回调函数
	fmt.Println("这是第二个类的第二个接口回调函数call2", Phone2{member_id: 44, member_balance: 100, member_sex: true, member_nickname: "陈超"})
}

func main() {
	var phone Phone

	//	先示例第一个接口
	phone = new(Phone1)
	phone.call()
	phone.call2()

	//	示例第二个接口
	phone = new(Phone2)
	phone.call()
	phone.call2()

} */