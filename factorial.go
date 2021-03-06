// Go Program to Find Factorial of a Number
// Write a Go Program to Find factorial of a Number using For loop. The for loop (for i := 1; i <= factorialnum; i++) iteration starts at one and ends at user given value.
// Within the loop (factorial = factorial * i),
// we multiply each value with a factorial variable value.

package main

import "fmt"

func canlcFact(num int) int {
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	return factorial
}

func main() {
	var num int
	fmt.Println("enter no : ")
	fmt.Scanln(&num)

	fmt.Println(canlcFact(num))
}
