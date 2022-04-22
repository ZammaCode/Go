package main

import "fmt"

func main() {

	//declare and initialize an array
	arr := [5]int{1, 2, 3, 4, 5}

	//print length of anr
	fmt.Println("Length of arr =", len(arr))
	//print capacity of arr
	fmt.Println("Capacity of arr =", cap(arr))
}
