// Go Program to Calculate Profit or Loss
// This go program uses the actual product cost and the sales amount to calculate profit or loss. It uses an else if statement to print output.
// If the product cost is greater than the sales price, then loss.
// If the sales price is higher than the product cost, then the product is profit—otherwise, no profit or loss.

package main

import "fmt"

var pcost float64

func ProfitLoss(pcost float64, scost float64) string {
	if pcost > scost {
		fmt.Println("You are in loss !!! ")
	} else {
		fmt.Println("You are in Profit !")
	}
	return ""
}

func production(pnum float64, pvalue float64) float64 {
	pcost = pnum * pvalue
	return pcost
}

func main() {
	var scost, pnum, pvalue float64

	fmt.Println("Enter Total product : ")
	fmt.Scanln(&pnum)

	fmt.Println("Enter product value : ")
	fmt.Scanln(&pvalue)

	fmt.Println("Enter sales cost : ")
	fmt.Scanln(&scost)

	fmt.Println("Production cost is : ", production(pnum, pvalue))

	fmt.Println(ProfitLoss(pcost, scost))
}
