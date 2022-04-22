package main

import (
	"fmt"
)

func main() {
	//2D array decleration and initialization
	//with 2 rows and 3 columns
	table := [2][3]string{{"A", "B", "C"}, {"D", "E", "F"}}
	//traversing 2D array using for-range loop
	for _, row := range table {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}

	fmt.Println("-----------------------")
	//3D array decleration and initialization
	//with 4 rows, 3 columns and depth of 2
	cube := [4][3][2]int{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{7, 8},
			{9, 10},
			{11, 12},
		},
		{
			{13, 14},
			{15, 16},
			{17, 18},
		},
		{
			{19, 20},
			{21, 22},
			{23, 24},
		},
	}
	//traversing 3D array using for loop
	for row := 0; row < 4; row++ {
		for col := 0; col < 3; col++ {
			for depth := 0; depth < 2; depth++ {
				fmt.Printf("%d \t", cube[row][col][depth])
			}
			fmt.Println()
		}
		fmt.Println("..........")
	}
}
