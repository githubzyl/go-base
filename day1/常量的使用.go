package main

import "fmt"

func main(){
	//常量的声明关键字：const

	const a int = 10
	fmt.Println(a)

	//常量的自动推导类型，不需要:=
	const b = 11.2
	fmt.Printf("b type is %T\n", b)
	fmt.Println(b)

}
