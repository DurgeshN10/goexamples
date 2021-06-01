// Go Program to calculate Compound Interest
// Write a Go Program to calculate Compound Interest. This golang program allows the user to enter the principal amount, totals years, and interest rates and then find the Compound Interest.

// Future Compound Interest =  principal amount * (1 + interest rates)years
// Compound Interest =  Future Compound Interest – principal amount

package main

import (
	"fmt"
	"math"
)

func main() {
	var cifuture, p, r, t, compoundI float64

	fmt.Println("Enter principle amout : ")
	fmt.Scanln(&p)

	fmt.Println("enter interest rate : ")
	fmt.Scanln(&r)

	fmt.Println("enter investment years : ")
	fmt.Scanln(&t)

	cifuture = p * (math.Pow((1 + r/100), t))
	compoundI = cifuture - p

	fmt.Println("future interest : ", cifuture)
	fmt.Println("Compound interest : ", compoundI)

}
