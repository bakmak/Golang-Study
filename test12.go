/* 
package main

import "fmt"

func main() {
	// 下面实例演示了创建和使用map:
	// var countryCapitalMp map[string]string	//	创建集合
	// countryCapitalMp = make(map[string]string)

	// // map插入key - value对，各个国家对应的首都
	// countryCapitalMp ["France"] = "巴黎"
	// countryCapitalMp ["Italy"] = "罗马"
	// countryCapitalMp ["Japan"] = "东京"
	// countryCapitalMp ["India"] = "新德里"

	// for cuntry := range countryCapitalMp {		//	使用键输出地图值
	// 	fmt.Println(cuntry,"首都是",countryCapitalMp[cuntry])
	// }

	// capital, ok := countryCapitalMp ["American"]			//		查看元素在集合中是否存在
	// fmt.Println("1",capital)
	// fmt.Println("2",ok)
	// if (ok) {		//		如果确定是真实的,则存在,否则不存在
	// 	fmt.Println("American 的首都是",capital)
	// } else {
	// 	fmt.Println("American 的首都不存在")
	// }
	
	// delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：
	countryCapitalMp := map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New delhi"}
	fmt.Println("原始地图")

	for countr := range countryCapitalMp {						//	打印地图
		fmt.Println(countr,"首都是",countryCapitalMp[countr])
	}

	delete(countryCapitalMp,"France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	for country := range countryCapitalMp {
		fmt.Println(country,"首都是",countryCapitalMp[country])
	}
} */