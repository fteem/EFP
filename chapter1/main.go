package main

import (
	"fmt"
)

func main() {
	var bill = 0
	var tipPercentage = 0

	for {
		fmt.Println("What is the bill?")
		fmt.Scanf("%d", &bill)

		if bill > 0 {
			break
		}
		fmt.Println("Please enter a valid number for the bill amount.")
	}

	for {
		fmt.Println("What is the tip percentage?")
		fmt.Scanf("%d", &tipPercentage)

		if tipPercentage > 0 {
			break
		}
		fmt.Println("Please enter a valid number for the tip percentage amount.")
	}

	fmt.Printf("The tip is $%d.\n", calculateTip(bill, tipPercentage))
	fmt.Printf("The total is $%d.\n", calculateTip(bill, tipPercentage)+bill)
}

func calculateTip(bill int, percentage int) int {
	return bill * percentage / 100
}
