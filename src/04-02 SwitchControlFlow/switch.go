package main

import (
	//time package provides functionality
	//for measuring and displaying time
	"fmt"
	"time"
)

func main() {
	//Now returns the current date and time
	//Weekday returns day of week
	today := time.Now().Weekday()

	//switch keyword is followed by test
	//condition
	switch today {
	case time.Monday:
		//if the condition specified with switch
		//statement matches Monday then this block
		//will be executed
		fmt.Println("Today is Monday")
	case time.Tuesday, time.Wednesday:
		//case is defined with more than one values
		//in this case if condition is matched with
		//either Tuesday or Wednesday then this block
		//will be executed
		if today == time.Wednesday {
			//if today is Wednesday then using the break
			//statement escape the current case block
			//the remaning code in the block will be not
			//executed
			fmt.Println("Today is Wednesday")
			break
		}
		fmt.Println("Today is either Tuesday or Wednesday")
	default:
		//if switch condition dose not match any of the case
		//specified above then default block will be executed
		fmt.Println("No match found")
	}

	fmt.Println("--------fallthrough demonstration----------")

	//declare switch condition with initialization statement
	switch i := 3; {
	case i == 3:
		//because i==3 condition is true therefore this block
		//will be executed
		fmt.Println("i==3 case evaluated to true")
		//the fallthrough statement will force
		//the execution flow to fall to the next case block
		fallthrough
	case i == 2:
		//due to fallthrough this block
		//will be executed despite that condition is false
		fmt.Println("i==2 case executed")
		fallthrough
	case i == 1:
		//due to fallthrough this block will
		//also be executed
		fmt.Println("i==1 case executed")
	}

	fmt.Println("--------type switch demonstration----------")
	//declare variable i with data type as interface
	var i interface{}

	//assign value to i
	i = 78

	//in siwtch condition we use the type assertion expression
	//it assert the type of value or variable
	//insted of usig specific data type we use type keyword
	//to be generic for all data types
	switch i.(type) {
	case string:
		fmt.Println("i is of type string")
	case bool:
		fmt.Println("i is of type bool")
	case int:
		fmt.Println("i is of type int")
	}
}
