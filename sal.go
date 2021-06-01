// Go Program to Calculate Employee Salary
// This Go program uses the Else If statement to calculate the employee’s gross salary. Please replace the HRA and DA percentages as you need.
package main

import "fmt"

func main() {
	var grosSal, basicSal, hra, da float64
	fmt.Println("Enter basic salary of Employee : ")
	fmt.Scanln(&basicSal)

	if basicSal <= 10000 {
		hra = basicSal * 8 / 100
		da = basicSal * 10 / 100
	} else if basicSal <= 20000 {
		hra = basicSal * 10 / 100
		da = basicSal * 12 / 100
	} else if basicSal <= 35000 {
		hra = basicSal * 12 / 100
		da = basicSal * 16 / 100
	}
	grosSal = basicSal + hra + da
	fmt.Println("Gross salary of employee : ", grosSal)
}
