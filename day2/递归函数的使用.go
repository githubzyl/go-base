package main

import "fmt"

func main(){
	sum := sum2(100)
	fmt.Println("sum = ", sum)
}

func sum1() (b int){
	for i:= 1; i<=100; i++ {
		b += i
	}
	return
}
func sum2(a int) int{
	if a <= 1 {
		return 1
	}
	return a + sum2(a-1)
}