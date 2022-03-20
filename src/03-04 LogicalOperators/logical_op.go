package main

import "fmt"

func main() {
	var (
		a    int = 6
		b    int = 10
		rslt bool
	)

	//logical AND operator
	rslt = a > 5 && b < 10
	fmt.Println("a>5 && b<10 ", rslt)

	//logical OR operator
	rslt = a == 5 || b <= 10
	fmt.Println("a==5 || b<=10 ", rslt)

	//logical NOT operator
	rslt = !(a == 6)
	fmt.Println("!(a==6) ", rslt)
}
