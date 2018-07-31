package main

import "fmt"

func main(){

	a := 10
	str := "mike"

	//匿名函数，没有函数名字，函数定义，还没有被调用
	f1 := func(){
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	//执行函数
	f1()

	//给一个函数类型起名
	type FuncType func()
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func(){
		fmt.Printf("a = %d, str = %s\n", a , str)
	}()//后面的()代表调用此匿名函数

	//带参数的匿名函数
	func(i, j int){
		fmt.Printf("i = %d, j = %d\n", i , j)
	}(1, 2)

	//匿名函数，有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)

	fmt.Printf("x = %d, y = %d\n", x, y)

	//闭包捕获外部变量的特点
	func(){
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a , str)
	}()
	fmt.Printf("外部：a = %d, str = %s\n", a , str)

}
