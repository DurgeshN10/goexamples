// Go Program to Calculate Cube of a Number
// This go program to calculate the cube of a number, we used arithmetic operator (multiplication) to find the cube.

package main

import "fmt"

func cube(num int) int {
	return num * num * num
}

func main() {
	var num int

	fmt.Println("Enter number : ")
	fmt.Scanln(&num)

	fmt.Println(cube(num))
}
