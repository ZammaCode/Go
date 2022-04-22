package main

import "fmt"

func main() {
	//declare array n which will hold 10 integer elements
	var arr1 [10]int

	var i int
	//array traversal using the for loop
	for i = 0; i < 10; i++ {
		//set element at index i to value of i
		arr1[i] = i
		//output each array element's value
		fmt.Printf("arr1[%d] = %d \n", i, arr1[i])
	}

	fmt.Println("-------------------------------")

	arr2 := [3]string{"Andriod", "Windows", "iOs"}
	//array traversal using the for-range loop,
	//range keyword returns index and value
	//j will hold the index of the array
	//v will hold the value of an element/item of array
	for j, v := range arr2 {
		fmt.Printf("Index=%d, value=%s \n", j, v)
	}

	//we can use a blank identifier to skip the index or element
	for _, v := range arr2 {
		fmt.Printf("Value=%s \n", v)
	}
}
