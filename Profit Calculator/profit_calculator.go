package main

import (
	"fmt"
	"os"
)

var taxRate float64 = 18
func main() {

	revenue := getUserInput("Revenue is: ")
	expenses := getUserInput("Expenses were: ")
	taxRate = getUserInput("Tax rate: ")

	// earningBeforeTax := revenue - expenses
	// profit := earningBeforeTax * (1 - taxRate/100)
	// ratio := earningBeforeTax / profit

	earningBeforeTax, profit, ratio := calculator(revenue, expenses)

	fmt.Println("EBT is ", earningBeforeTax, ", Profit is ", profit, " and Ratio was ", ratio)
	results := fmt.Sprintf("EBT is %0.2f, Profit is %0.4f, and Ratio was %0.4f", earningBeforeTax, profit, ratio)
	fmt.Printf(results)


	os.WriteFile("Metrics.txt", []byte(results), 0644)
}

func calculator(revenue, expenses float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoTest string) float64 {
	var inputValue float64
	fmt.Print(infoTest)
	fmt.Scan(&inputValue)
	if inputValue <= 0 {
		panic("Value should be greater than 0")
	}
	return inputValue
}
