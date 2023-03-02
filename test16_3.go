/* package main

import "fmt"

type DIV_ERR struct {		//	自定义错误信息结构
	etype int 				//	错误类型
	v1 int					//	记录下出错时的陈数、被除数
	v2 int
}

func (div_err DIV_ERR) Error() string {		//	实现接口方法	error.Error()
	if 0 == div_err.etype {
		return "除零错误"
	} else {
		return "其他未知错误"
	}
}

func div(a int, b int) (int, *DIV_ERR) {		//	除法
	if b == 0 {									//	返回错误信息
		return 0, &DIV_ERR{0,a,b}
	} else {
		return a / b, nil
	}
}

func main() {
	v, r := div(100, 2)					//	正确调用
	if nil != r {
		fmt.Println("(1)fail:",r)
	} else {
		fmt.Println("(1)succeed:",v)
	}

	v, r = div(100, 0)					//	错误调用
	if nil != r {
		fmt.Println("(2)fail:", r)
	} else {
		fmt.Println("(2)succeed:", v)
	}
}

 */