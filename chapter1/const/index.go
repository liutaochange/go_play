package main

import "fmt"

func main()  {
	fmt.Println("const")
	// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	// const identifier [type] = value
	// 你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	// 显式类型定义： const b string = "abc"
	// 隐式类型定义： const b = "abc"
	// 多个相同类型的声明可以简写为：
	// const c_name1, c_name2 = value1, value2
	// 常量还可以用作枚举：
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	// 数字 0、1 和 2 分别代表未知性别、女性和男性。
	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过

	// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	// 简写方式
	const (
		x = iota
		y
		z
	)
	fmt.Println(x, y, z)

	const (
		e = iota   //0
		g          //1
		f          //2
		d = "ha"   //独立值，iota += 1
		h          //"ha"   iota += 1
		k = 100    //iota +=1
		l          //100  iota +=1
		m = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,g,f,d,h,k,l,m,i)


	const (
		v=1<<iota
		w=3<<iota
		q
		r
	)
	// v=1：左移 0 位,不变仍为 1;
	// w=3：左移 1 位,变为二进制 110, 即 6;
	// q=3：左移 2 位,变为二进制 1100, 即 12;
	// r=3：左移 3 位,变为二进制 11000,即 24。
	fmt.Println("v=",v)
    fmt.Println("w=",w)
    fmt.Println("q=",q)
	fmt.Println("r=",r)
	
	// 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
	const (
		im = 0
		ig
		iq
	)
	fmt.Println(im, ig, iq) // 0 0 0 
}