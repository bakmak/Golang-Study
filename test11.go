/* package main

import (
	"fmt"
	"os"
)

// var pow = []int{1,2,4,8,16,32,64,128}
func f0()  {		//		Go Range 简单循环：
	nums := []int{1,2,3,4}
	length := 0
	for range nums {
		length++
	}
	fmt.Println(length)
}

func f1()  {		//	循环键值对
	nums := []int{1,2,3,4}
	for i, num := range nums {
		fmt.Printf("索引是%d, 长度是%d\n", i, num)
	}
}

func f2()  {		//	通过 range 获取参数列表:
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

func printStr(s string)  {			//	Go 中的中文采用 UTF-8 编码，因此逐个遍历字符时必须采用 for-each 形式：
	fmt.Println("str: " + s)
	for _, v := range s {
		fmt.Printf("0x%x %c\t", v,v)
	}
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x\t", s[i])
	}

}

func f3()  {
	printStr("hello")
	fmt.Println("")
	printStr("中国人")
}

func main() {
	// f0()
	// f1()
	// f2()
	f3()

	// 遍历简单的数组，2**%d 的结果为索引对应的次方数：
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// for 循环的 range 格式可以省略 key 和 value，如下实例：
	// map1 := make(map[int]float32)
	// map1[1] = 1.0
	// map1[2] = 2.0
	// map1[3] = 3.0
	// map1[4] = 4.0

	// for key, value := range map1 {					//		读取key和value
	// 	fmt.Printf("key is : %d - value is : %f\n", key,value)
	// }

	// for key := range map1 {							//		读取 key
	// 	fmt.Printf("key is : %d\n", key)
	// }

	// for _, value := range map1 {					//		读取value
	// 	fmt.Printf("value is : %f\n", value)
	// }

	//	range 遍历其他数据结构：
	// nums := []int{2,3,4}	//	这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	// sum := 0
	// for _, num := range nums {	
	// 	sum += num
	// }
	// fmt.Println("sum:",sum)

//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用了空白符"_"省略了。有时候我们确实需要知道它的索引
	// for i, num := range nums {
	// 	if num == 3 {
	// 		fmt.Println("index:",i)
	// 	}
	// }

	// kvs := map[string]string{"a":"apple","b":"banana"}		//	range 也可以用在 map 的键值对上
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }

	// //	range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身
	// for i, c := range "go" {								
	// 	fmt.Println(i, c)
	// }

}



 */