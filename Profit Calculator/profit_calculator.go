package main

import "fmt"

func main() {
	const taxRate float64 = 18
	var revenue, expenses float64

	fmt.Print("Enter Revenue Generated: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit

	fmt.Println("EBT is ", earningBeforeTax, ", Profit is ", profit, " and Ratio was ", ratio)
}
