/* package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"bytes"
)

func test1()  {
	f4, _ := os.Open("a.txt")
	defer f4.Close()
	r := bufio.NewReader(f4)
	s, _ := r.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2()  {
	r := strings.NewReader("ABCEFG")
	r2 := strings.NewReader("123456")
	r3 := bufio.NewReader(r)
	s, _ := r3.ReadString('\n')
	fmt.Println(s)
	r3.Reset(r2)
	s2, _ := r3.ReadString('\n')
	fmt.Println(s2)
}

func test3()  {
	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	r2 := bufio.NewReader(r)
	b := make([]byte, 10)

	for {
		n, err := r2.Read(b)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(b[0:n]): %v\n", string(b[0:n]))
		}
	}
}

func test4()  {
	r := strings.NewReader("ABGEFG")
	r2 := bufio.NewReader(r)

	b, _ := r2.ReadByte()
	fmt.Printf("b: %c\n", b)

	b2, _ := r2.ReadByte()
	fmt.Printf("b2: %c\n", b2)

	r2.UnreadByte()
	b3, _ := r2.ReadByte()
	fmt.Printf("b3: %c\n", b3)
}

func test5()  {
	r := strings.NewReader("你好，世界")
	r2 := bufio.NewReader(r)

	r3, size, _ := r2.ReadRune()
	fmt.Printf("r3: %c, size: %v\n", r3, size)

	r4, size2, _ := r2.ReadRune()
	fmt.Printf("r4: %c, size2: %v\n", r4, size2)

	r2.UnreadRune()
	r5, size3, _ := r2.ReadRune()
	fmt.Printf("r5: %c, size3: %v\n", r5, size3)
}

func test6()  {
	r := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	r2 := bufio.NewReader(r)

	line, isPrefix, _ := r2.ReadLine()
	fmt.Printf("line: %q, isPrefix: %v\n", line, isPrefix)

	line2, isPrefix2, _ := r2.ReadLine()
	fmt.Printf("line2: %q, isPrefix2: %v\n", line2, isPrefix2)

	line3, isPrefix3, _ := r2.ReadLine()
	fmt.Printf("line3: %q, isPrefix3: %v\n", line3, isPrefix3)

	line4, isPrefix4, _ := r2.ReadLine()
	fmt.Printf("line4: %q, isPrefix4: %v\n", line4, isPrefix4)
}

func test7()  {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
}

func test8()  {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
}

func test10()  {
	f4, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer f4.Close()
	w := bufio.NewWriter(f4)
	w.WriteString("ijnefb")
	w.Flush()

}

func test9()  {
	s := strings.NewReader("ABCEFGHIJKLMN")
	br := bufio.NewReader(s)
	// b := bytes.NewBuffer(make([]byte, 0))
	b, _ := os.OpenFile("a.txt", os.O_RDWR, 0775)
	defer b.Close()
	br.WriteTo(b)
	// fmt.Printf("%s\n", b)
}

func test11()  {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)       
	fmt.Println(c) 
}

func test12()  {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available()) 
	fmt.Println(bw.Buffered())  
	fmt.Printf("%q\n", b)      

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())  
	fmt.Printf("%q\n", b)    
}

func test13()  {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	// 写入缓存
	// byte等同于 int8
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	// rune等同于int32
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('！')
	// 写入b
	bw.Flush()
	fmt.Println(b)
}

func test14()  {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 世界！")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	//bw.Flush()            //ReadFrom无需使用Flush，其自己已经写入．
	fmt.Println(b) // Hello 世界！
}

func test15()  {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p))              //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b)
}

func test16()  {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func test17()  {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}

func test18()  {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	} 
}

func main() {
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	// test14()
	// test15()
	// test16()
	// test17()
	// test18()
} */