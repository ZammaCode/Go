package main

import "fmt"

func main() {
	//declare two constants of type string
	const NAME, GENDER string = "Asghar Khan", "Male"

	//print value of constants using the formating verbs
	//here %v means print in default format
	//\n is used for newline
	fmt.Printf("Value of NAME= %v, and GENDER= %v \n", NAME, GENDER)

	//this statement will generate error as
	//we cant change value of constant
	//NAME = "Khan"
	//fmt.Println(NAME)

	//declare untyped const
	const FLAG = true

	fmt.Printf("Value of FLAG= %v \n", FLAG)
}
