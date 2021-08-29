package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

func main() {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1) // 10

	fmt.Println("=====分隔线=====")

	// t2 := i.(string)
	// fmt.Println(t2) // panic

	// other version
	// var j interface{} //nil
	// t3 := j.(interface{})
	// fmt.Println(t3) // panic 因为nil会触发panic

	t1, err := i.(int)
	fmt.Printf("%d-%t\n", t1, err) //10-true

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok) //<nil> - false

	findType(10) // int
	//测试
}
