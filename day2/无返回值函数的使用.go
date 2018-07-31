package main

import "fmt"

func main(){
	//Test()
	//Test21(1,2)
	//Test3(1,2,3)
	//Test31(1, 2,3,4)
	Test4(1,2,3,4)
}

//无参无返回的函数
func Test(){
	fmt.Println("无参无返回的函数")
}

//有参无返回的函数，普通参数列表
func Test2(a int, b int){
	fmt.Println("a + b = ", a + b)
}
func Test21(a , b int){
	fmt.Println("a + b = ", a + b)
}
func Test22(a int, b string, c float64){
	fmt.Println(a, b, c)
}

//不定参数
func Test3(args ...int){//传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args))

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("=================================")

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}
//固定参数+不定参数
func Test31(a int, args ...int){
	fmt.Println("a = ", a)
	fmt.Println("agrs = ", args)
}

//不定参数在多个函数中传递
func Test4(args ...int){
	//全部参数传递
	//Test41(args...)

	Test42(args[:2]...) //从参数下标0到2的值传递到其他函数（不包括下标为2的值）
	fmt.Println("===========")
	Test42(args[2:]...) //从参数下标2到最后的值传递到其他函数（包括下标为2的值）
}
func Test41(tmp ...int){
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}
func Test42(tmp ...int){
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}