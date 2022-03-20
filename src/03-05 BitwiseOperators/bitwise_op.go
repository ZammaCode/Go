package main

import (
	"fmt"
	//strconv package is used to convert data types to a string format
	//and also from a string to other basics data types
	"strconv"
)

func main() {
	var a int64 = 11
	var b int64 = 12
	var rslt int64

	//strconv.FormatInt(number,base) converts
	//integer to binary equivalent
	fmt.Println("11 in binary = ", strconv.FormatInt(a, 2))
	fmt.Println("12 in binary = ", strconv.FormatInt(b, 2))
	fmt.Println("-----------------------")

	//Bitwise AND operation
	rslt = a & b
	fmt.Println("11 & 1011 = ", strconv.FormatInt(rslt, 2))

	//Bitwise OR operation
	rslt = a | b
	fmt.Println("11 | 1011 = ", strconv.FormatInt(rslt, 2))

	//Bitwise XOR operation
	rslt = a ^ b
	fmt.Println("11 ^ 1011 = ", strconv.FormatInt(rslt, 2))

	//Bitwise left shift operation
	//shift by two bits to the left
	rslt = a << 2
	fmt.Println("11 << 2 = ", strconv.FormatInt(rslt, 2))

	//Bitwise right shift operation
	//shift by 3 bits to the right
	rslt = a >> 3
	fmt.Println("11 >> 3 = ", strconv.FormatInt(rslt, 2))
}
