package main

import "fmt"

//函数的返回值是一个匿名函数，返回一个函数类型
func test() func() int {
	var x int //没有初始化， 值为0
	return func() int {
		x++
		return x * x
	}
}

func main(){

	//返回值是一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，也就是f来调用闭包函数
	//它不关心这些捕获了的变量和常量是否超出了作用域，所以只要闭包还在使用它，这些变量就存在
	f := test()

	fmt.Println(f())//1
	fmt.Println(f())//4
	fmt.Println(f())//9
	fmt.Println(f())//16
	fmt.Println(f())//25

}
