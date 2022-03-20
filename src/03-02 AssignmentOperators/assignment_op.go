package main

import "fmt"

func main() {
	//declare variables
	var a int

	//assigns right side value to the variable on the left
	a = 5

	//declare variable c and assign value 0
	//and also associate respective data type
	c := 0
	fmt.Println("c := 0", c)
	//add, subtract, multiplay, divide
	//values of variables and also assign to the left variable
	a += 10
	fmt.Println("a += 10", a)
	a -= 2
	fmt.Println("a -= 2", a)
	a *= 5
	fmt.Println("a *= 5", a)
	a /= 10
	fmt.Println("a /= 10", a)
	a %= 2
	fmt.Println("a %= 2", a)

}
