/* package main

import "fmt"

func f()  {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			break
		}
		fmt.Printf("i : %v\n", i)
	}
}

func f1()  {
	i := 2
	switch i {
	case 1:
		fmt.Println("等于1")
		break
	case 2:
		fmt.Println("等于2")
		// break
		fallthrough
	case 3:
		fmt.Println("等于3")
		break
	default:
		fmt.Println("不关心")
		break
	}
}

func f2()  {
MY_LABEL:
	for i := 0; i < 10; i++ {
		if i >= 5 {
			break MY_LABEL
		}
		fmt.Printf("%v\n",i)
	}
	fmt.Println("end...")
}


func main() {
	// f()
	// f1()
	f2()
} */