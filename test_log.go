/* package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init()  {
	f4, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
        log.Panic("打开日志文件异常")
    }
	logger = log.New(f4, "success", log.Ldate | log.Ltime | log.Lshortfile)

}

func testLog_print_panic_fatal()  {
	defer fmt.Println("发生了 panic错误！")
	log.Print("my log")
	log.Printf("my log %d",100)
	name := "tom"
	age := 20
	log.Println( name, "," , age)
	log.Fatal("致命错误！")
	fmt.Println("end...")

}

func testLog_Flags_SetFlags()  {
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("my log")
}

func testLog_Prefix_SetPrefix()  {
	s := log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.SetPrefix("MyLog: ")
	s = log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.Print("my log...")
}

func testLogOutput()  {
	f4, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Print("打开日志文件异常")
	}
	log.SetOutput(f4)
	log.Print("my log...")
}

func main() {
	// testLog_print_panic_fatal()
	// testLog_Flags_SetFlags()
	// testLog_Prefix_SetPrefix()
	// testLogOutput()
	logger.Println("自定义logger")
} */