package main

import "fmt"

func main() {
	//outer loop
	for i := 0; i < 3; i++ {
		//inner loop
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d \n", i, j)

			if i == j {
				fmt.Println("Breaking out of inner loop")
				//break here from the inner loop
				//and continue with outer loop
				break
			}
		}
	}

	fmt.Println("--------------------")
	//to stop printing when both i and j are equal
	//we will use label with break statement

	//define label
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d \n", i, j)
			if i == j {
				fmt.Println("Breaking out of inner loop")
				//exit from the inner loop
				//and the execution flow will continue after
				//outer label
				break outer
			}
		}
	}
}
