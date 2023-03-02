/* 
package main

import (
	// "fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// r := strings.NewReader("hello world")
	// buf := make([]byte, 20)
	// r.Read(buf)
	// fmt.Printf("string(buf): %v\n", string(buf))

	// r := strings.NewReader("some io.Reader stream to be read\n")
	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 		log.Fatal(err)
	// }

	r := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second read\n")
	buf := make([]byte, 8)

	//	Buf用在这里…
	if _, err := io.CopyBuffer(os.Stdout, r, buf); err != nil {
		log.Fatal(err)
	}

	//	．.． 重用也在这里。不需要分配额外的缓冲区。
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

} */