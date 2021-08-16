package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

func Print(obj interface{}) error {
	//取Value
	v := reflect.ValueOf(obj)

	for i := 0; i < v.NumField(); i++ {
		// 去里面的每个领域
		// v.Type().Field(i).Tag 取i的Tag
		// v.Field(i) 取i的值
		field := v.Type().Field(i)
		tag := field.Tag

		label := tag.Get("label")
		value := fmt.Sprintf("%v", v.Field(i))

		if value == "" {
			value = tag.Get("default")
		}
		fmt.Println(label + value)
	}
	return nil
}

func main() {
	p1 := Person{
		name: "yxf",
		Age:  18,
	}
	Print(p1)
}
