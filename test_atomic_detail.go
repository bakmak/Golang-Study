/* package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub()  {		//	加减
	var i int32 = 100
	atomic.AddInt32(&i,1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i,-1)
	fmt.Printf("i: %v\n", i)

	var j int64 = 200
	atomic.AddInt64(&j,1)
	fmt.Printf("j: %v\n", j)
}

func test_load_stora()  {		//	读和写
	var i int32 = 100
	atomic.LoadInt32(&i)		//	read
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i,200)
	fmt.Printf("i: %v\n", i)	//	write
}

func test_cas()  {				//	直接交换
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)	//	i的100换为200
	fmt.Printf("b: %v\n", b)	//	交换成功就会返回true，错误就是flash
	fmt.Printf("i: %v\n", i)	
}

func main() {
	// test_add_sub()
	// test_load_stora()
	test_cas()
} */