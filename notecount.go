// Go Program to Count Total Notes in an Amount
// Write a Go Program to Count the total number of currency notes in a given amount using Arrays and For loop. First, we declared an integer array that holds the available notes.
// Next, we used for loop (for i := 0; i < 8; i++) to iterate the notes array and divides the amount with each array item.
// Then, we update the count by removing that cash amount from the original.

// package main

// import "fmt"

// func main() {
// 	notes := [8]int{200, 100, 50, 20, 10, 5, 2, 1}
// 	var amount int

// 	fmt.Println("Enter amount : ")
// 	fmt.Scanln(&amount)

// 	temp := amount
// 	for i := 0; i < 8; i++ {
// 		fmt.Println(notes[i], "notes = ", temp/notes[i])
// 		temp = temp % notes[i]
// 	}
// }
package main

import "fmt"

func noteCount(amount int) int {
	notes := [8]int{200, 100, 50, 20, 10, 5, 2, 1}
	temp := amount

	for i := 0; i < 8; i++ {
		fmt.Println(notes[i], "notes = ", temp/notes[i])
		temp = temp % notes[i]
	}
	return amount
}

func main() {
	var amount int

	fmt.Println("Enter amount : ")
	fmt.Scanln(&amount)

	fmt.Println(noteCount(amount))
}
