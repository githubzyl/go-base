package main

import "fmt"

func main(){
	var a int
	fmt.Println("请输入变量a:")

	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a)//&很重要
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
