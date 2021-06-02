// Go Program to find Generic Root of a Number
// Write a Go Program to find the generic root of a Number using For loop. The generic root of a number is the sum of all digits until the sum is less than 10.
// For instance, generic root of 98 = 9 + 8 = 17 => 1 + 7 = 8. To achieve this, we used nested for loop (for sum = 0; genNumber > 0; genNumber = genNumber / 10) to find the sum of digits.
// And the if else statement (if sum >= 10) to check the sum greater than or equal to ten further divides the value.

package main

import "fmt"

func root(num int) int {
	var sum, reminder int

	for sum = 0; num > 0; num = num / 10 {
		reminder = num % 10
		sum = sum + reminder
	}
	return sum
}

func main() {
	var num, sum int

	fmt.Println("enter number : ")
	fmt.Scanln(&num)
	for num >= 10 {
		sum = root(num)
		if sum >= 10 {
			num = sum
		} else {
			fmt.Println("Root is : ", sum)
			break
		}
	}

}
