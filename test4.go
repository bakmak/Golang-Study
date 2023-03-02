/* 
package main

import "fmt"

func f0()  {	//	输出1-100素数
	var C,c int	//	声明变量
	C = 1	//	这里不写入for循环是因为for语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）
	A : for C < 100 {
		C++	// C = 1不能写入for这里就不能写入
		for c = 2; c < C; c++ {
			if C % c == 0 {
				goto A		//	若发现因子不是素数
			}
		}
		fmt.Println(C,"是素数")
	}
}

func f2()  {
	var a , b int
	for a = 2; a <= 100; a++ {
		for b = 2; b <= (a / b); b++ {
			if a % b == 0 {
				break
			}
		}
		if b > (a / b) {
			fmt.Println(a,"是素数")
		}
	}
}

func f1()  {	//	九九乘法表
	for i := 1; i <= 9; i++ {	//	i 控制行，以及计算的最大值
		for j := 1; j <= i; j++ {	//	j 控制每行的计算个数
			fmt.Printf("%d * %d = %d\t", j , i , i*j)
		}
		fmt.Println("")
	}
}

func main() {
	// f2()
	// f1()
	// f0()		
	// //		无限循环
	// for true {
	// 	fmt.Printf("这是无限循环。\n")
	// }
}
  */
package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	var score int = 0
	var fenshu string = "A"
	fmt.Println("欢迎进入成绩查询系统")
ZHU:for true {
		var xuanzhe int = 0
		fmt.Println("1.进入成绩查询\t2.退出程序")
		fmt.Printf("请输入序号选择：")
		fmt.Scanln(&xuanzhe)
		fmt.Println("")
		if xuanzhe == 1 {
			goto CHA
		}
		if xuanzhe == 2 {
			os.Exit(1)
		}
	}

CHA:for true {
		fmt.Printf("请输入一个学生的成绩：")
		fmt.Scanln(&score)

		switch  {
			case score >= 90:	fenshu = "A"
			case score >= 80 && score < 90:	fenshu = "B"
			case score >= 60 && score < 80:	fenshu = "C"
			default: fenshu = "D"
		}

		var c string = strconv.Itoa(score)
		switch {
			case fenshu == "A":
				fmt.Printf("你考了%s分，评价为%s，成绩优秀\n", c,fenshu)
			case fenshu == "B" || fenshu == "C":
				fmt.Printf("你考了%s分，评价为%s，成绩良好\n", c,fenshu)
			case fenshu == "D":	
				fmt.Printf("你考了%s分，评价为%s，成绩不合格\n", c,fenshu)
		}
		fmt.Println("")
		goto ZHU
	}
}





