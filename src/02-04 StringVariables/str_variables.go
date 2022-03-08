package main

import "fmt"

func main() {
	//declare hw variable statically using
	//data type string and also assign a value
	var hw string = "Hello World"

	fmt.Println(hw)

	//declare msg variable dynamically using :=
	//operator and make it multiline by using backtick marks
	msg := `This is 
			an example of multi-line
			string.`

	fmt.Println(msg)
}
