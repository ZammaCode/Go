package main

import "fmt"

func main() {
	//outer loop will iterate for number of lines
	//that's  it will print 5 number of lines
	for i := 0; i < 5; i++ {
		//inner loop will iterate for number of .
		//that is it will print 4 number of .
		//for each iteration of outer loop
		for j := 0; j <= 3; j++ {
			fmt.Print(".")
		}

		//print new line character
		fmt.Print("\n")
	}
}
