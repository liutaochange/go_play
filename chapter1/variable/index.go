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

	// 如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。
	// 如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误 a declared and not used。
	// 全局变量是允许声明但不使用。 

	var a string = "go"
    fmt.Println(a)

    var b, c int = 1, 2
	fmt.Println(b, c)

	print()

	// 多变量可以在同一行进行赋值
	x, y, z := 5, 7, "str"  // 同时赋值未声明变量
	fmt.Println(x,y,z)

	// 同时赋值已声明变量
	var e int
	var s string
	var h bool
	e, s, h = 10, "study", true
	fmt.Println(e, s, h)

	// 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
	x, e = e, x
	fmt.Println(x, e) // 10, 5
	// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	// _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	_, h = true, false
	// fmt.Println(_, h) // cannot use _ as value
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