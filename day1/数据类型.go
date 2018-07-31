package main

import "fmt"

func main(){
	//bool类型
	var a = true
	fmt.Println("a = ", a)

	//2、浮点型
	var b = 3.24
	fmt.Println("b = ", b)

	//3、字符类型byte
	var c byte
	c = 97
	fmt.Printf("%d, %c\n", c, c)

	c = 'a'
	fmt.Printf("%d, %c\n", c, c)

	//4、字符串类型
	var str1 string;
    str1 = "hello go"
    fmt.Println("str1 = ", str1)
    //内建函数，len
    fmt.Println(len(str1), '\n')
    fmt.Printf("str1[0] = %c\n", str1[0])

    //5、复数类型
    var cp complex128
    cp = 2.1 + 3.14i
    fmt.Println("cp = ", cp)
    t := 2.1 + 3.14i
    fmt.Printf("t type is %T\n", t)
    //通过内建函数，获取实部和虚部
    fmt.Println("real(t) = ", real(t), ", imag(t) = ", imag(t))



}
