package main

import "fmt"

func main() {
	//declare array n which will hold 10 integer elements
	var n [10]int

	var i int
	for i = 0; i < 10; i++ {
		//set element at index i to value of i
		n[i] = i
		//output each array element's value
		fmt.Printf("n[%d] = %d \n", i, n[i])
	}

	fmt.Println("=============================")

	//declare array arr of type string with 5 elements
	//and initialize it with literal values
	arr := [5]string{"AB", "CD", "EF", "GH", "IJ"}
	//accessing individual elements of any array
	fmt.Printf("arr[2] = %s \n", arr[2])
}
