/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strings"
)

func testReadAll()  {
	// r := strings.NewReader("hello world!")		//	实现字符串

	f4, err := os.Open("a.txt")						//	实现文件
	defer f4.Close()								//	执行之后要把文件关掉

	b, err := ioutil.ReadAll(f4)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func testReadDir()  {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func testReadFile()  {
	b, _ := ioutil.ReadFile("a.txt")
	fmt.Printf("string(b): %v\n", string(b))
}

func testWriteFile()  {
	ioutil.WriteFile("a.txt",[]byte("golang"), 00664)
}

func testTempFile()  {
	b := []byte("temporary file's content")
	f4, err := ioutil.TempFile("", "bakmak")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("f4.Name(): %v\n", f4.Name())
	defer	os.Remove(f4.Name())		//	clean up

	if _, err := f4.Write(b); err != nil {
		log.Fatal(err)
	}
	if err := f4.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	// testReadDir()
	// testReadFile()
	// testWriteFile()
	testTempFile()
} */