/* package main

import "fmt"

func f5()  {
	score := 80
	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score >= 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
	if score := 100; score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score >= 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func f6()  {
	// Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	var c string
	fmt.Println("输入一个字符：")
	fmt.Scan(&c)

	if c == "S" {
		if c == "a" {
			fmt.Println("Saturday")
		} else if c == "u" {
			fmt.Println("Sunday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("输入第二个字符")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else if c == "h" {
			fmt.Println("Thursday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else {
		fmt.Println("输入错误")
	}



}

func main() {
	// f5()
	f6()
} */