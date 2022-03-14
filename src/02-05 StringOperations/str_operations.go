package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//declare and initialize name variable statically
	var name string = "Asghar Khan"

	/*with %v access modifier value at position 2
	will be printed in default format which in the case
	of strings is byte. to print in character format we
	can use %c modifier*/
	fmt.Printf("Character at position 2 of slice in byte format is %v \n", name[1])
	fmt.Printf("Character at position 2 of slice in character format is %c \n", name[1])

	msg := "This is hello world! message"
	//Split method slices a msg into substrings seperated by !
	//the seperator e.g. ! is not included in the substring
	sp := strings.Split(msg, "!")
	//print substring slices
	fmt.Println(sp[0])
	fmt.Println(sp[1])

	//SplitAfter method slices a msg into substrings seperated by !
	//and the seperator e.g. ! is included in the substring
	sp = strings.SplitAfter(msg, "!")
	fmt.Println(sp[0])
	fmt.Println(sp[1])

	//hw := "Hello"
	//assign character L at index 0 of hw string
	//when we run this line of code compiler will
	//generate an error
	//hw[0] = 'L'

	str1 := "Hello"
	str2 := "World"
	//concatenate  str1 and str2 using + operator
	result := str1 + " " + str2
	fmt.Println(result)

	//compare str1 and str2
	var rslt bool
	//if both strings are equal then true will
	//be returned and assigned to rslt otherwise false
	rslt = str1 == str2
	fmt.Println(rslt)

	//declare and assign unicode string str3
	//which will hold two urdu words
	str3 := "محمد اصغر"
	//the len function will return the no of
	//bytes used to store two urdu words
	fmt.Println(len(str3))
	//the RuneCountInString will return actual
	//numbers of characters of both words
	fmt.Println(utf8.RuneCountInString(str3))

}
