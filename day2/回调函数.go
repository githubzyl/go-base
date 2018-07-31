package main

import "fmt"

type FunType func(int, int) int

func Add1(a, b int) int {
	return a + b
}

func Minus1(a, b int) int {
	return a - b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，实现四则运算
func Calc(a, b int, fTest FunType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	return
}

func main(){
	a := Calc(1, 1, Add1)
	fmt.Println("a = ", a)
	a = Calc(1, 1, Minus1)
	fmt.Println("a = ", a)
}
