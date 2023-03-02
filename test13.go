/* package main



import (
	"fmt"
	"unsafe"
)

//	递归，就是在运行的过程中调用自己。
//	语法格式如下：
func recursion()  {
	recursion()
}

//	阶乘		以下实例通过 Go 语言的递归函数实例阶乘：
func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n - 1)
		return result
	}
	return 1 
}

//	斐波那契数列	以下实例通过 Go 语言的递归函数实现斐波那契数列：
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

//	更好的一种 fibonacci 实现，用到多返回值特性，降低复杂度：
func fibonacci_0(n int) (int,int) {
	if n < 2 {
		return 0, n
	}
	a,b := fibonacci_0(n - 1)
	return b,a+b
}
func fibonacci_1(n int) int {
	_,b := fibonacci_0(n)
	return b
}

// 原理: 计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：
func sqrt(x float64 , i float64) (float64,float64) {
	remain := (i * i - x) / (2 * i)
	i = i - remain
	if (remain > 0) {
		return sqrt(x,i)
	} else {
		return i, remain
	}
}
// func get_sqrt(x float64) float64 {
// 	i,_ := sqrt(x,x)
// 	return i
// }

func get_sqrt(x float32) float32 {
	xhalf := 0.5*x
	var i int32 = *(*int32)(unsafe.Pointer(&x))
	i = 0x5f375a86 - (i >> 1)
	x = *(*float32)(unsafe.Pointer(&i))

	x = x * (1.5 - xhalf * x * x)
	x = x * (1.5 - xhalf * x * x)
	x = x * (1.5 - xhalf * x * x)

	return 1/x
}

func main() {
	// var i int = 3
	// fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	// recursion()
	
	// var i int
	// for i = 0; i < 10; i++ {
	// 	fmt.Printf("i:%d\t", fibonacci_1(i))
	// }

	// fmt.Println(get_sqrt(2))
	// fmt.Println(get_sqrt(3))
	
	fmt.Println(get_sqrt(4))
}





 */