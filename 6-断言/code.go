package main

import "fmt"

func main() {
	i := 10
	fmt.Print(i.(int))
	fmt.Print(i)
}
