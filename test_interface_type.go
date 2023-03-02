/* package main

import "fmt"

//	一个类型实现多个接口
// type Player interface {			//	定义一个Player接口
// 	playMusic()
// }

// type Video interface {			//	定义一个Video接口
// 	playVideo()
// }

// type Mobile struct {			//	定义Mobile结构体
// }

// //	实现两个接口
// func (m Mobile) playMusic() {	
// 	fmt.Println("播放音乐")
// }

// func (m Mobile) playVideo() {
// 	fmt.Println("播放视频")
// }

//	多个类型实现同一个接口
type Pet interface {				//	定义一个Pet接口
	eat()
}

type Dog struct {					//	定义一个Dog结构体
}

type Cat struct {					//	定义一个Cat结构体
}

//	实现接口
func (cat Cat) eat()  {				
	fmt.Println("cat eat...")
}

func (dog Dog) eat()  {
	fmt.Println("dog eat...")
}

func main() {
	// m := Mobile()
	// m.playMusic()
	// m.playVideo()

	var p Pet		//	声明一个Pet接口
	p = Cat{}
	p.eat()
	p = Dog{}
	p.eat()
} */