// Go Program to Count Digits in a Number
// Write a Go Program to Count Digits in a Number using for loop. The for loop condition (num > 0) make sure the number is greater than zero.
// Within the loop, we are incrementing the count value. Next, we divide the number by ten (num = num / 10), which will remove the last digit from a number.

// package main

// import "fmt"

// func main() {
// 	var num, count int //float64

// 	fmt.Println("enter number : ")
// 	fmt.Scanln(&num)

// 	for num > 0 {
// 		num = num / 10
// 		count = count + 1
// 	}
// 	fmt.Println("number of digits in number is : ", count)
// }

// Golang Program to Count Digits in a Number
// In this Golang example, we altered the for loop and counted the total number of individual digits in a user-given value.
// for count = 0; num > 0; num = num / 10 {
// 	count = count + 1
// }

package main

import "fmt"

var count int = 0

func digitCount(num int) int {
	if num > 0 {
		count = count + 1
		digitCount(num / 10)
	}
	return count
}

func main() {

	var num, count int
	fmt.Println("Enter digit : ")
	fmt.Scanln(&num)

	count = digitCount(num)

	fmt.Println("number of digit : ", count)
}
