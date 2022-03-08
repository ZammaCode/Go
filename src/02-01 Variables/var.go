/*declare package as main to make our
module as exe*/

package main

import (
	//import the format package from Go standard library
	"fmt"
	//import the reflect package from  Go standard library
	"reflect"
)

//decalre main function
func main() {

	/*declare variables statically using var keyword
	declare age variable with a data type of int
	and also assigns value*/
	var age int8 = 35

	fmt.Println("## Variable declared using var keyword:")
	fmt.Println("Value of age", age)
	fmt.Println("Data Type of age", reflect.TypeOf(age))

	fmt.Println("-------------------------------------")

	/*declare variables dynamically using := operator
	variable hw will be declared with data type inferred
	from  the value assigned which will be string in this case*/
	hw := "Muhammad Asghar Khan"

	fmt.Println("## Variable declared using := operator")
	fmt.Println("Value of hw", hw)
	fmt.Println("Data Type of hw", reflect.TypeOf(hw))

	fmt.Println("-------------------------------------")

	//declare multipale variables in a block
	var (
		flag1 bool = true
		temp  = 28.4
	)

	fmt.Println("Variable declared in block")
	fmt.Println("Value of flag1", flag1)
	fmt.Println("Type of flag1", reflect.TypeOf(flag1))
	fmt.Println("Value of temp", temp)
	fmt.Println("Type of temp", reflect.TypeOf(temp))

}
