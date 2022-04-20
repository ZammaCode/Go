package main

import "fmt"

func main() {

	i := 1
	for i < 5 {
		//when i is equal to 2, increment the value of i by 1
		if i == 2 {
			i++
			//skip this iteration and jump to start of loop
			continue
		}
		fmt.Println(i)
		i++
	}
}
