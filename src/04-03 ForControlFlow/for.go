package main

import "fmt"

func main() {

	//for loop traditional form example
	fmt.Println("---- Demonstraton of for traditional form ----")

	/*the i:=0 is initialization statement which executes
	before start of loop
	the condition statement is always a boolean expression which
	evaluates at each iteration of loop, if condition is true then
	loop will execute
	the update statement i++ will be executed after each iteration,
	after which condition is evaluated again*/
	for i := 0; i <= 2; i++ {
		fmt.Println("Value of i = ", i)
	}

	//for loop in the form of while loop
	fmt.Println("---- Demonstraton of for loop while form ----")

	//declare and initialize loop variable
	j := 0
	//itrate till condition evaluates to true
	for j < 2 {
		fmt.Println("Value of j = ", j)

		//update or increment loop variable
		j++
	}

	//for range form
	fmt.Println("---- Demonstraton of for range form ----")

	//declare and initialize string variable
	str := "Hi"

	//with the range keyword for loop iterate over
	//a collection of items (e.g. in this case string)
	//and assign both its index and value to loop variables
	//e.g. range stores character and index of str on each
	//itration
	for idx, chr := range str {
		fmt.Printf("The index number of %q is %v \n", chr, idx)
	}
}
