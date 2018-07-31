package main

import "fmt"

func main(){

	s := 10
	if s > 15 {
		fmt.Println(s > 5)
	}else if s > 9 {
		fmt.Println(s > 9)
	}else{
		fmt.Println(s == 10)
	}

	//if支持一个初始化语句,初始化语句和判断条件以分号分隔
	if a := 11; a == 10 {
		fmt.Println("a == 10 :", a == 10)
	}else {
		fmt.Println("a != 10 :", a != 10)
	}

	//switch支持一个初始化语句,初始化语句和判断条件以分号分隔
	switch num :=5 ; num {
		case 1:
			fmt.Println("F1")
			//break
			//fallthrough
		case 2:
			fmt.Println("F2")
			//break
		case 3,4,5:
			fmt.Println("F345")
		//break
		default:
			fmt.Println("xxx")
	}
	
	score := 12
	switch {//可以没条件
	case score > 90:
		fmt.Println("A")
	case score > 80:
	    fmt.Println("B")
	case score > 70:
		fmt.Println("C")
	case score > 60:
		fmt.Println("D")
	default:
		fmt.Println("不及格")
	}

	//for循环，循环只有这一种，没有别的
	//1+2+3+..1100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum+=i
		if sum <= 50 {
			continue
		}else if sum >= 4000 {
			break
		}
	}
	fmt.Println("累加结果为：", sum)

	//range迭代
	str := "abcdefg"
	//通过for打印每个字符
	/*
	for i :=0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	*/
	//迭代打印每个元素
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

}
