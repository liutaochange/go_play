package main

import "fmt"

func main()  {
	// Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
	// 声明变量的一般形式是使用 var 关键字：
	// 	var identifier type
	// 可以一次声明多个变量：
	// 	var identifier1, identifier2 type

	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	// var v_name v_type
	// v_name = value
	// 零值就是变量没有做初始化时系统默认设置的值

	var a string = "go"
    fmt.Println(a)

    var b, c int = 1, 2
	fmt.Println(b, c)

	print()

}


// 数值类型（包括complex64/128）为 0
// 布尔类型为 false
// 字符串为 ""（空字符串）
func print()  {

	// 没有初始化就为空
	var a string
	fmt.Println(a)

	// 没有初始化就为零值
    var b int
    fmt.Println(b)

    // bool 零值为 false
    var c bool
    fmt.Println(c)
}