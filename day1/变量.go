package main

import "fmt"

//go函数可以返回多个值
func test()(a, b, c int){
	return 1, 2, 3
}

func main(){

	//1、声明格式：var 变量名 类型,变量声明了必须要用
	//2、只是声明没有初始化的变量，默认值为0
	//3、同一个{}，声明的变量名是唯一的
	var a int
	fmt.Println("a=",a)

	//4、可以同时声明多个变量
	var b, c int
	fmt.Println("b=", b)
	c=10
	fmt.Println("c=", c)

	//5、变量初始化，声明变量时同时赋值
	var d int = 10
	d = 20
	fmt.Println("d=",d)

	//5、自动推导类型，必须初始化，通过初始化的值确定类型(非常实用)
	//:=，先声明e的类型，再赋值
	e := 30
	//%T打印变量所属的类型
	fmt.Printf("c type is %T\n", e)

	//6、多重赋值，匿名变量
    f, g := 10, 20
    fmt.Printf("f = %d, g = %d\n", f, g)

    //交换两个变量的值
    var tmp int
    tmp = f
    f = g
    g = tmp
	fmt.Printf("f = %d, g = %d\n", f, g)

    //交换两个变量的值（so easy）
    f, g =  g, f
	fmt.Printf("f = %d, g = %d\n", f, g)

    //_匿名变量，丢弃不处理;
    tmp, _ = g, f
	fmt.Printf("tmp = %d, g = %d\n", tmp, g)

    var h, i, j int
    h, i, j = test();
    fmt.Printf("h = %d, i = %d, j = %d\n", h, i, j)

    //_匿名变量配合函数返回值使用，才有优势
    _, i, j = test()
    fmt.Printf("i = %d\n", i)

}
