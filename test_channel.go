/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)		

func send()  {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)			//	发送后等待5秒钟才接收到消息
	values <- value
}

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <- values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")
} */