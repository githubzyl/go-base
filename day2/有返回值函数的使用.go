package main

import "fmt"

func main(){
	//fmt.Println("Test1() = ", Test1())
	//fmt.Println("Test11() = ", Test11())
	//fmt.Println("Test12() = ", Test12())

	//a,b,c := Test13()
	a,b,c := Test14()
	fmt.Printf("a=%d, b=%d, c=%d\n ", a,b,c)

	max, min := Test15(12, 5)
	fmt.Printf("max=%d, min=%d\n ", max , min)

	//通过匿名变量丢弃某个返回值
	max1, _ := Test15(12, 5)
	fmt.Printf("max1=%d\n ", max1)
}

//一个返回值
func Test1() int {
	return 666
}
//给返回值定义一个变量名，go推荐写法
func Test11() (result int) {
	return 666
}
//常用写法
func Test12() (result int) {
	result = 666
	return
}

//多个返回值
func Test13() (int, int, int){
	return 1,2,3
}
//go官方推荐写法
func Test14() (a, b, c int){
	 a,b,c = 1,2,3
	return
}

//有参数有返回值
func Test15(a,b int) (max, min int) {
	if a >= b {
		max, min = a, b
	}else{
		max, min = b, a
	}
	return
}