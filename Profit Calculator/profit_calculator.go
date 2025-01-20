package main

import "fmt"

const taxRate float64 = 18
func main() {

	var revenue, expenses float64

	fmt.Print("Enter Revenue Generated: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	// earningBeforeTax := revenue - expenses
	// profit := earningBeforeTax * (1 - taxRate/100)
	// ratio := earningBeforeTax / profit

	earningBeforeTax, profit, ratio := calculator(revenue, expenses)


	fmt.Println("EBT is ", earningBeforeTax, ", Profit is ", profit, " and Ratio was ", ratio)
	fmt.Printf("EBT is %0.2f, Profit is %0.4f, and Ratio was %0.4f", earningBeforeTax, profit,  ratio)
}

func calculator(revenue, expenses float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
