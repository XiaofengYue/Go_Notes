package main

import (
	"fmt"
)

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)
}

// 输出
// myfunc 函数里的name：go
// main 函数里的name: python
// main 函数里的myname:  go
