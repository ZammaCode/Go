package main

import "fmt"

func main() {
	//declare numeric variables
	var (
		a int = 100
		b int = 20
		c int
	)

	//declare string variables
	var (
		str1 string = "Go"
		str2 string = "lang"
		str  string
	)

	//addition operator operation on integer data type
	c = a + b
	fmt.Println("a + b = ", c)
	//addition operator operation on string data type
	str = str1 + str2
	fmt.Println("str1 + str2 = ", str)
	//subtraction operator
	c = a - b
	fmt.Println("a - b = ", c)
	//multiplication operator
	c = a * b
	fmt.Println("a * b = ", c)
	//division operator
	c = a / b
	fmt.Println("a / b = ", c)
	//modulus operator
	c = a % b
	fmt.Println("a % b =", c)
	//increment operator
	a++
	fmt.Println("a++", a)
	//decrement operator
	a--
	fmt.Println("a--", a)

}
