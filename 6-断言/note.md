# 理解
检查 i 是否为 nil
检查 i 存储的值是否为某个类型

# Tips
i.(int)
i.(string)

但是又更加高级的不需要一个一个来判别
```
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
findType(10)      // int
findType("hello") // string
}
```

# 注意
```
t, ok:= i.(T)
```
和上面一样，这个表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回其类型给 t，并且此时 ok 的值 为 true，表示断言成功。

如果接口值的类型，并不是我们所断言的 T，就会断言失败，但和第一种表达式不同的事，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败，此时t 为 T 的零值。

# 什么时候使用
用来诊断错误的时候