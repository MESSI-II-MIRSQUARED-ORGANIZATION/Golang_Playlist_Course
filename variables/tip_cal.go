package main

import (
	"fmt"
)

func main(){
	const defaultTipPercentage float64 = 15.0

	var billAmount, tipPercentage float64

	// fmt.Println("Enter the bill amount: ")
	fmt.Scanln(&billAmount)


	fmt.Println("Enter the tip percentage (or press Enter to use default percentage)")
	n,_ := fmt.Scanln(&tipPercentage)

	if n == 0{
		tipPercentage = defaultTipPercentage
	}


	tip := (billAmount * tipPercentage) / 100

	fmt.Printf("The tip is: $%.2f\n", tip)
}