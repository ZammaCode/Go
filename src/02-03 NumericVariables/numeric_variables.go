/*declare package as main to make our module as exe*/

package main

import "fmt"

/*import the format package from Go standard library*/

//decalre main function
func main() {

    //declare variables statically using var keyword
    //declare age variable with a data type of int
    //and also assigns value
    var age int8 = 35

    //declare height variable with a data type of float
    //and also assign value
    var height float32 = 6.3

    fmt.Println("## Numeric Variables declared using var keyword:")
    fmt.Println(age)
    fmt.Println(height)

    //declare variables dynamically using := operator
    //variable x will be declared with data type inferred
    //from the value assigned which will be int in this case
    x := 2788887

    //variable y will be declared with data type float
    y := 787.100

    fmt.Println("## Numeric Variables declared using := operator")
    fmt.Println(x)
    fmt.Println(y)
}
