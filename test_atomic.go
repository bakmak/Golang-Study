
//	原子操作
/* package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add()  {
	atomic.AddInt32(&i,1)
}

func sub()  {
	atomic.AddInt32(&i,-1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)
} */


//	使用锁实现协程的同步
/* package main

import (
	"fmt"
	"sync"
	"time"
)

var i = 100

var lock sync.Mutex

func add()  {
	lock.Lock()
	i++
	lock.Unlock()
}

func sub()  {
	lock.Lock()
	i--
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)
} */

