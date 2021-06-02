// Go Program to Check Leap year
// A leap year contains 366 days, and any year divisible by four except the century years are called leap years.
// Furthermore, if the century year divisible by 400, then it is a leap year.
// In this Go program, we use the If else statement (if yr%400 == 0 || (yr%4 == 0 && yr%100 != 0)) and logical operators to check leap year or not.

package main

import "fmt"

func main() {
	var yr int

	fmt.Println("Enter year : ")
	fmt.Scanln(&yr)

	if yr%400 == 0 || (yr%4 == 0 && yr%100 != 0) {
		fmt.Println(yr, "Is a leap year ")
	} else {
		fmt.Println("entered year is not leap year")
	}
}
