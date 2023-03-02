/* package main

import (
	"fmt"
	"time"
)

func show(msg string)  {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go show("java")
	show("golang")			//	在main协程中执行，如果它前面也添加go，程序没有输出
	fmt.Println("end...")
} */