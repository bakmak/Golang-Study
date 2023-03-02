/* package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func testTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)
	fmt.Printf("b: %T\ti: %T\n", b, i)
	fmt.Printf("b: %v\ti: %v\n", b, i)
	fmt.Println("----------------------------------")
	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
	fmt.Printf("s: %T\tb1: %T\n", s, b1)
	fmt.Printf("s: %v\tb1: %v\n", s, b1)
}


func testContains()  {
	b := []byte("duoke360.com")
	sublice1 := []byte("duoke360.com")
	sublice2 := []byte("DUOKE360.com")
	fmt.Printf("bytes.Contains(b, sublice1): %v\n", bytes.Contains(b, sublice1))
	fmt.Printf("bytes.Contains(b, sublice2): %v\n", bytes.Contains(b, sublice2))
	fmt.Printf("strings.Contains(\"hello wrold\", \"hello\"): %v\n", strings.Contains("hello wrold", "hello"))
}

func testCount()  {
	b := []byte("helloooooooooooooooooooooooooo")
	b2 := []byte("h")
	b3 := []byte("l")
	b4 := []byte("o")
	fmt.Printf("bytes.Count(b2): %v\n", bytes.Count(b,b2))
	fmt.Printf("bytes.Count(b3): %v\n", bytes.Count(b,b3))
	fmt.Printf("bytes.Count(b4): %v\n", bytes.Count(b,b4))
}

func testRepeat()  {
	b := []byte("hi")
	fmt.Printf("string(bytes.Repeat(b, 1)): %v\n", string(bytes.Repeat(b, 1)))
	fmt.Printf("string(bytes.Repeat(b, 3)): %v\n", string(bytes.Repeat(b, 3)))
}

func testReplace()  {
	b := []byte("hello,wrold")
	b2 := []byte("o")
	b3 := []byte("ee")
	fmt.Printf("string(bytes.Replace(b, b2, b3, 0)): %v\n", string(bytes.Replace(b, b2, b3, 0)))
	fmt.Printf("string(bytes.Replace(b, b2, b3, 1)): %v\n", string(bytes.Replace(b, b2, b3, 1)))
	fmt.Printf("string(bytes.Replace(b, b2, b3, 2)): %v\n", string(bytes.Replace(b, b2, b3, 2)))
	fmt.Printf("string(bytes.Replace(b, b2, b3, -1)): %v\n", string(bytes.Replace(b, b2, b3, -1)))
}

func testRunes()  {
	s := []byte("巴卡玛卡 巴卡玛卡")
	r := bytes.Runes(s)
	fmt.Printf("转换前字符串的长度：%v\n", len(s))
	fmt.Printf("转换后字符串的长度：%v\n", len(r))
}

func testJoin()  {
	b := [][]byte{[]byte("你好"), []byte("世界")}
	b2 := []byte(",")
	fmt.Printf("string(bytes.Join(b, b2)): %v\n", string(bytes.Join(b, b2)))
	b3 := []byte("#")
	fmt.Printf("string(bytes.Join(b, b3)): %v\n", string(bytes.Join(b, b3)))
}

func testReader() {
	data := "123456789"
	//	通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))
	//	返回未读取部分的长度
	fmt.Printf("re.Len(): %v\n", re.Len())
	//	返回底层数据总长度
	fmt.Printf("re.Size(): %v\n", re.Size())
	fmt.Println("-----------------------------")

	buf := make([]byte, 2)
	for {
		//	读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("-----------------------------")

	//	设置偏移量，因为上面的操作已经修改了读取位置等信息
	re.Seek(0, 0)
	for {
		//	一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("------------------------------")

	re.Seek(0,0)
	off := int64(0)
	for {
		//	指定编译量读取
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func testBuffer()  {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}

func testBuffer2()  {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b.Bytes(): %v\n", string(b.Bytes()))
}

func testBuffer3()  {
	var b = bytes.NewBufferString("hello")
	b1 := make([]byte,2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(b1[:n]): %v\n", string(b1[:n]))
	}
}

func main() {
	// testTrans()
	// testContains()
	// testCount()
	// testRepeat()
	// testReplace()
	// testRunes()
	// testJoin()
	// testReader()
	// testBuffer()
	// testBUffer2()
	testBuffer3()
} */