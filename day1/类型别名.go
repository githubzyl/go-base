package main

import "fmt"

func main(){

	type myint int

	var i myint
	i = 10
	fmt.Println("i = ", i)

	type(
		long int64
		char byte
	)

	var k long
	var c char

	k = 100
	c = 'a'
	fmt.Printf("k = %d, c = %c\n", k, c)

}
