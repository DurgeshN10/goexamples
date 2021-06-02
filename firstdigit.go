// Go Program to find First Digit of a Number
// Write a Go Program to find the First Digit of a Number using For loop. The for loop condition (for firstDigit >= 10) returns true until the number is greater than or equals to 10.
// We divide the number by 10 (firstDigit = firstDigit / 10), which will return the first value after all the iterations.

package main

import "fmt"

func firstdig(num int) int {
	for num >= 10 {
		num = num / 10
	}
	return num
}

func main() {
	var num int
	fmt.Println("enter digit : ")
	fmt.Scanln(&num)

	fmt.Println(firstdig(num))
}
