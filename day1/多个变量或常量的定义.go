package main

import "fmt"

func main(){

	//多个变量不同类型的同时声明
	var(
		a int = 2
		b float64 = 2.1
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	const(
		i int = 20
		j float64 = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
