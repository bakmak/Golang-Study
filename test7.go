/* package main

import "fmt"

func f0()  {
	var n [10]int	//	n 是一个长度为10的数组
	var i, j int

	for i = 0; i < 10; i++ {	//	为数组 n 初始化元素
		n[i] = i + 100 			//	设置元素为i + 100
	}

	for j = 0; j < 10; j++ {	//	输出
		fmt.Printf("Elemen[%d] = %d\n", j, n[j])
	}
}

func f1()  {
	var i, j, k int
	balance := [5]float32{1000.0,2.0,3.4,7.0,50.0}		//	声明数组的同时快速初始化数组

	for i = 0; i < 5; i++ {		//	输出数组元素
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}
	fmt.Println("")
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {	//	输出每个数组元素的值
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}
	fmt.Println("")
	balance3 := [5]float32{1:2.0,3:7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

func GetYangHuiTriangleNextLine(inArr []int) []int  {
	var out []int
	var i int
	arrLen := len(inArr)
	out = append(out,1)
	if 0 == arrLen {
		return out
	}
	for i = 0; i < arrLen-1; i++ {
		out = append(out, inArr[i] + inArr[i+1])
	}
	out = append(out,1)
	return out
}

func triangle(n int)  {
	var item []int
	for i := 1; i <= n; i++ {
		item_len := len(item)
		if item_len == 0 {
			item = append(item,1)
		} else {
			temp_s := []int{1}
			for j := 0; j < item_len-1; j++ {
				temp_s = append(temp_s,item[j]+item[j+1])
			}
			temp_s = append(temp_s,1)
			item = temp_s
		}
		fmt.Println(item)
	}
}

func yanghuisanjiao(rows int) {
	for i := 0; i < rows; i++ {
		number := 1
		for k := 0; k < rows-i; k++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%5d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println("")
	}
}

func f2()  {
	var value = [3][2]int{{1,2},{3,4},{5,6}}	//	二维数组
	for i, v := range value {					//	遍历二维数组的其他方法，使用 range
		for j, v2 := range v {
			fmt.Printf("value[%v][%v] = %v \t", i, j, v2)
		}
		fmt.Print(v)
		fmt.Println("")
	}
}

func main() {
	// f0()
	// f1()

	// nums := []int{}
	// var i int
	// for i = 0; i < 10; i++ {
	// 	nums = GetYangHuiTriangleNextLine(nums)
	// 	fmt.Println(nums)
	// }

	// triangle(13)

	// yanghuisanjiao(20)

	// f2()
} */