/* package main

import (
	"fmt"
	"strconv"
)

const LINES int = 10	//	行数

func ShowYangHuiTriangle()  {		//		杨辉三角形
	nums := []int{}
	for i := 0; i < LINES; i++ {
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length - i] + nums[length - i - 1]
			}
			nums = append(nums,value)
			fmt.Print(value," ")
		}
		fmt.Println("")
	}
}

//九九乘法表 
func add(a,b int) int {
	return a + b
}
func multiplicationTable()  {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			var ret string
			if i * j < 10 && j != 1 {
				ret = " " + strconv.Itoa(i*j)
			} else {
				ret = strconv.Itoa(i * j)
			}
			fmt.Print(j," * ",i," = ",ret,"    ")
		}
		fmt.Printf("\n")
	}
}

func max(num1, num2 int) int {		//	函数调用	函数返回两个数的最大值
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x,y string) (string,string)  {	//函数返回多个值
	return y,x
}

func linxin()  {	//	返回菱形
	x := 30 			//	长
	y := 30			//	宽
	row := 1		//	行数
	for row <= y {
		count := 0
		if row <= (y/2 + 1) {
			count = (2 * row - 1)
		} else {
			count = 2 * (y - row) + 1
		}
		row++
		text := ""
		star_min := ((x - count) / 2) + 1	//	算出显示星星的范围
		star_max := ((x - count) / 2) + count
		for index := 1; index <= x; index++ {
			if index >= star_min && index <= star_max {
				text += "*"
			} else {
				text += " "
			}
		}
		fmt.Println(text)
	}
}

func multiplicationTable_silce()  {		//	切片解决九久乘法表
	var num []int 	//定义切片
	for i := 1; i < 10; i++ {
		num = append(num, i)	//将数据添加到切片中去
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < i+ 1; j++ {
			value := num[j-1]*i		//	计算
			fmt.Printf("%d*%d=%d\t", j,i,value)
		}
		fmt.Println("")
	}
}

func main() {
	// var a int = 200
	// var b int = 100
	// var ret int

	// ret = max(a,b)		//	调用函数并返回最大值

	// fmt.Printf("最大值是 ： %d\n", ret)

	// a , b := swap("Google","Runoob")
	// fmt.Println(a,b)
	// ShowYangHuiTriangle()
	// multiplicationTable()
	// linxin()
	// multiplicationTable_silce()
}


 */