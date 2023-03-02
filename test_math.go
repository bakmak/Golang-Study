/* package main

import (
	"math/rand"
	"fmt"
	"math"
	"time"
)

func test1()  {
	fmt.Printf("math.MaxFloat64: %.f\n", math.MaxFloat64)
	fmt.Printf("math.SmallestNonzeroFloat64: %.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("math.MaxFloat32: %.f\n", math.MaxFloat32)
	fmt.Printf("math.SmallestNonzeroFloat32: %.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("math.MaxInt8: %v\n", math.MaxInt8)
	fmt.Printf("math.MaxUint8: %v\n", math.MaxUint8)
	fmt.Printf("math.MinInt8: %v\n", math.MinInt8)
	fmt.Printf("math.MaxInt16: %v\n", math.MaxInt16)
	fmt.Printf("math.MinInt16: %v\n", math.MinInt16)
	fmt.Printf("math.MaxUint16: %v\n", math.MaxUint16)
	fmt.Printf("math.MaxInt32: %v\n", math.MaxInt32)
	fmt.Printf("math.MinInt32: %v\n", math.MinInt32)
	fmt.Printf("math.MaxUint32: %v\n", math.MaxUint32)
	fmt.Printf("math.MaxInt64: %v\n", math.MaxInt64)
	fmt.Printf("math.MinInt64: %v\n", math.MinInt64)
	fmt.Printf("math.Pi: %.2f\n", math.Pi)
}

func test2()  {
	//	取绝对值，函数签名如下
	//	func Abs(x float64)	float64		
	fmt.Printf("[-3.14]绝对值为:[%.2f]\n", math.Abs(-3.14))

	//	取x的y次方，函数签名如下：
	//	func Pow(x, y float64) float64	
	fmt.Printf("[2]的16次方为: [%.f]\n", math.Pow(2, 16))

	//	取余数，函数签名如下：
		// func Pow10(n int) float32	
	fmt.Printf("10的[3]次方为: [%.f]\n", math.Pow10(3))

	//	取x的开平方，函数签名如下：
		// func Sqrt(x float64) float64	
	fmt.Printf("[64]的开平方为: [%.f]\n", math.Sqrt(64))

	//	取x的开立方，函数签名如下：
		// func Cbrt(x float64) float64	
	fmt.Printf("[27]的开立方为: [%.f]\n", math.Cbrt(27))

	//	向上取整，函数签名如下：
		// func Ceil(x float64) float64	
	fmt.Printf("[3.14]向上取整为: [%.f]\n", math.Ceil(3.14))

	//	向下取整，函数签名如下：
		// func Floor(x float64) float64	
	fmt.Printf("[8.75]向下取整为: [%.f]\n", math.Floor(8.75))

	//	取余数，函数签名如下：
		// func Mod(x, y float64) float64	
	fmt.Printf("[10/3]的余数为: [%.f]\n", math.Mod(10, 3))

	//	分别取整数和小数部分，函数签名如下：
		// func Modf(f float64) (int float64, frac float64)	
	Integer, Decimal := math.Modf(3.14159265358979)
	fmt.Printf("[3.14159265358979]的整数部分为:[%.f],小数部分为:[%.14f]\n", Integer, Decimal)
}

func init() {
	//以时间作为初始化种子
	// rand.Seed(1)
	rand.Seed(time.Now().UnixNano())
}

func test3()  {
	// for i := 0; i < 10; i++ {	
	// 	i2 := rand.Int()
	// 	fmt.Println(i2)
	// }
	// fmt.Println("----------------------------")
	// for i := 0; i < 10; i++ {	//	范围
	// 	i2 := rand.Intn(100)
	// 	fmt.Println(i2)
	// }
	fmt.Println("------------------------------")
	for i := 0; i < 10; i++ {		//	小数
		f4 := rand.Float32()
		fmt.Println(f4)
	}
}

func main() {
	// test1()
	// test2()
	test3()
} */