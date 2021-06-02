// Go program to Find Power of a Number
// This Go program uses the math pow function to find the power of a number.
// To use this function, you have to import the math module.
package main

import (
	"fmt"
	"math"
)

func power(num float64, expo float64) float64 {
	power := math.Pow(num, float64(expo))
	return power
}

func main() {
	var num, expo float64

	fmt.Println("Enter number : ")
	fmt.Scanln(&num)

	fmt.Println("Enter exponent number : ")
	fmt.Scanln(&expo)

	fmt.Println("power of ", num, "is", power(float64(num), float64(expo)))
}
