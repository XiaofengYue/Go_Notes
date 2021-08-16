# 理解
\`json:"name"\`
``用作Tag的取值

# Tips
利用反射模式可以得到其中的Tag值和真正的值
`key01:"value01" key02:"value02" key03:"value03"`

# 注意

- json.Marsha() 
返回的是两个对象，用两个变量来接收
- v.Type().Field(i).Tag 
取i的Tag
- v.Field(i) 
取i的值
- reflect.ValueOf(obj) 
反射模式取得struct的值 因为这样能访问到Tag
- i < v.NumField()

# 什么时候使用
一般在需要格式化输出使使用。
Eg：改变输出格式