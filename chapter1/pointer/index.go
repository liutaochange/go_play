package main

import "fmt"

// 一个指针变量指向了一个值的内存地址。
// 类似于变量和常量，在使用指针前你需要声明指针。
// var var_name *var-type
// var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。

// 指针使用流程：
// 1. 定义指针变量。
// 2. 为指针变量赋值。
// 3. 访问指针变量中指向地址的值。
// 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

// go 空指针
// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
// nil 指针也称为空指针。
// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
// 一个指针变量通常缩写为 ptr。

func main()  {
	fmt.Println("Pointer")
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)
	
	var ip *int
    fmt.Printf("变量的地址：%x\n", &a)
    fmt.Println("变量的地址：", &a)
    ip = &a
    fmt.Println("ip 变量存储的指针地址:", ip)
    fmt.Println("ip 变量存储的指针地址的值:", *ip)
    fmt.Println("ip 变量存储的指针地址的地址:", &ip)
    var ptr *int
    if (ptr != nil) {
        if (ip != nil) {         
            fmt.Println("ptr不是空指针")     
            fmt.Println("ip不是空指针")   
        }else {       
            fmt.Println("ptr不是空指针")      
            fmt.Println("ip是空指针")     
        }   
    } else {  
        if(ip != nil){      
            fmt.Println("ptr是空指针")     
            fmt.Println("ip不是空指针")      
        }else{        
            fmt.Println("ptr是空指针")  
            fmt.Println("ip是空指针")    
		} 
	}
}