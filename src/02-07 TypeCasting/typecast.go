package main

import "fmt"

func main() {

	var f_var float32 = 1.5
	var i_var int8 = 2

	//try implicit casting by assigning
	//float to integer variable
	//on execution the following line will generate error
	//i_var = f_var

	//explcit type casting float into integer
	i_var = int8(f_var)
	fmt.Println(i_var)
}
