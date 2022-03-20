package main

import "fmt"

func main() {
	//decclare variables
	a := 10
	b := 20
	//declare c as boolean as comparision operators
	//return true and false
	var c bool

	c = a == b
	fmt.Println("a == b ", c)

	c = a != b
	fmt.Println("a != b ", c)

	c = a > b
	fmt.Println("a > b ", c)

	c = a >= b
	fmt.Println("a >= b ", c)

	c = a < b
	fmt.Println("a < b ", c)

	c = a <= b
	fmt.Println("a <= b ", c)

}
