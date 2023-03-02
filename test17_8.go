/* package main

import (
	"fmt"
	"time"
)

func put(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
		fmt.Println("->放入", i)
	}
	fmt.Println("=所有的都放进去了！关闭缓冲区，但是里面的数据不会丢失，还能取出。")
	close(c)	
}

func main() {
	ch := make(chan int, 5)
	go put(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		data, ok := <-ch
		if ok == true {
			fmt.Println("<-取出", data)
		} else {
			break
		}
	}	
} */