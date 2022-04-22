package main

import "fmt"

func main() {
	//declare array n which will hold 10 integer elements
	var n [10]int

	//assign value at index 2 of array n
	n[1] = 10
	//print value at specific index
	fmt.Println(n[1])

	//declare array m of type string with 5 elements
	//and initialize it with literal values
	m := [5]string{"AB", "CD", "EF", "GH", "IJ"}
	//print whole array m
	fmt.Println(m)

	//array initialization can ommit the size of array
	nums := [...]int{1, 2}
	fmt.Println(nums)
}
