/* package main

import "fmt"

//	组合接口
type reader interface {
	read() string
}

type writer interface {
	write() string
}

type rw interface {
	reader
	writer
}

type mouse struct {}

func (m mouse) read() string {
	return "mouse reading..."
}

func (m *mouse) write() string {
	return "mouse writing"
}

func main() {
	var rw1 rw
	rw1 = &mouse{}		//	只要有一个指针实现，则此处必须是指针

	fmt.Println(rw1.read())
	fmt.Println(rw1.write())
} */