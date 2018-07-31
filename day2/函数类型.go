package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给函数类型起名
//FuncType是一种函数类型
//入参类型及个数相同，返回类型相同
type FuncType func(int, int) int //没有函数名字，没有{}{

func main(){

	var result int
	var fTest FuncType
	fTest = Add
	result = fTest(20,10) //等价于Add(20,10)
	fmt.Println("result = ", result)

	fTest = Minus
	result = fTest(20,10) //等价于Minus(20,10)
	fmt.Println("result = ", result)

}
