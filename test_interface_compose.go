/* package main

import "fmt"

type Flyer interface {				//	定义一个飞Flyer接口
	fly()	
}

type Swimmer interface {			//	定义一个Swimmer接口
	swim()
}

type FlyFish interface {			//	将Flyer接口和Swimmer接口组合成FlyFish接口
	Flyer
	Swimmer
}

type Fish struct {					//	定义一个Fish结构体
}

//	实现这个组合接口
func (fish Fish) fly()  {	
	fmt.Println("fly...")
}

func (fish Fish) swim()  {
	fmt.Println("swim...")
}

func main() {
	var ff FlyFish		//	声明一个FlyFish类型接口变量
	ff = Fish{}			
	ff.fly()
	ff.swim()
} */