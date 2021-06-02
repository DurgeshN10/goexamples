// Go Program to Print Multiplication Table
// This Go program uses nested for loop to print multiplication tables of 9 and 10.
// The outer for loop to iterate values from 9 to 10, and the inner for loop to iterate from 1 to 20.
// If you want the multiplication table to display upto 20, change the nested for loop condition to j <= 20.

package main

import "fmt"

func main() {

	var i, j int

	fmt.Print("\nEnter the Number Less than or Equal to 10 = ")
	fmt.Scanln(&i)

	fmt.Println("\nMultiplication Tables upto 10 are = ")
	for i <= 10 {
		for j = 1; j <= 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
		i++
		fmt.Println("=======")
	}
}
