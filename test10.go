/* package main

import "fmt"

func prnitSlice(x []int)  {		//	len()和cap()函数
	fmt.Printf("len:%d cap:%d slice:%v\n",len(x),cap(x),x)
}

func f0()  {
	// 切片截取
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}	//	创建切片
	prnitSlice(numbers)

	fmt.Println("numbers == ",numbers)	//	打印原始切片

	fmt.Println("numbers[1:4] == ",numbers[1:4])	//	打印子切片从索引1（包含）到索引4（不包含）

	fmt.Println("numbers[:3] == ",numbers[:3])		//	默认下限为 0

	fmt.Println("numbers[4:] == ",numbers[4:])		//	默认上限为len(s)

	numbers1 := make([]int, 0, 5)
	prnitSlice(numbers1)

	numbers2 := numbers[:2]		//	打印子切片从索引 0(包含) 到索引 2(不包含)
	prnitSlice(numbers2)

	numbers3 := numbers[2:5]	//	打印子切片从索引 2(包含) 到索引 5(不包含)
	prnitSlice(numbers3)
}

func f1()  {
	var numbers []int
	prnitSlice(numbers)

	numbers = append(numbers, 0)	//	允许追加空切片
	prnitSlice(numbers)

	numbers = append(numbers, 1)	//	向切片添加一个元素
	prnitSlice(numbers)

	numbers = append(numbers, 2,3,4)	// 同时添加多个元素
	prnitSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)		//	创建切片 number1 是之前切片的两部容量

	copy(numbers1, numbers)		//	拷贝 numbers 的内容到 numbers1
	prnitSlice(numbers1)
}

func sliceTest()  {
	arr := []int{1,2,3,4,5}
	s := arr[:]
	for e := range s {
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}

func twoDimensionArray()  {
	var a = [][]int{{0,0},{1,2},{2},{3,6},{4,8}}	//	数组 - 5 行 2 列
	var i, j int

	for i = 0; i < len(a); i++ {					//	输出数组元素
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

}

func f2()  {			//	合并多个数组
	var arr1 = []int{1,2,3}
	var arr2 = []int{4,5,6}
	var arr3 = []int{7,8,9}
	var s1 = append(append(arr1,arr2...),arr3...)
	fmt.Printf("s1: %v\n", s1)

}

func main() {
	// var numbers []int

	// if (numbers == nil) {
	// 	fmt.Println("切片是空的")
	// }
	
	// f0()
	// f1()
	
	// sliceTest()
	// twoDimensionArray()

	f2()
}



 */