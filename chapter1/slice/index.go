package main

import "fmt"

// Go 语言切片是对数组的抽象。
// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，
// 可以追加元素，在追加时可能使切片的容量增大。

func main()  {
	fmt.Println("slice")
	// var identifier []type  切片不需要说明长度。
	var numbers = make([]int,3,5)
	fmt.Println(numbers);
	var slice = make([]int,5,100);
	fmt.Println(slice);

	// 也可以简写为
	// var slice2 := make([]type, len)
}