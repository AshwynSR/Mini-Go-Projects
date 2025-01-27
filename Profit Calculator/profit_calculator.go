package main

import (
	"fmt"
	"os"
)

const taxRate float64 = 18
func main() {

	var revenue, expenses float64

	fmt.Print("Enter Revenue Generated: ")
	fmt.Scan(&revenue)

	if revenue <= 0 {
		panic("Revenue can't be less then or equal to 0")
	}
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	if expenses <= 0 {
		panic("Expenses can't be less then or equal to 0")
	}

	// earningBeforeTax := revenue - expenses
	// profit := earningBeforeTax * (1 - taxRate/100)
	// ratio := earningBeforeTax / profit

	earningBeforeTax, profit, ratio := calculator(revenue, expenses)

	fmt.Println("EBT is ", earningBeforeTax, ", Profit is ", profit, " and Ratio was ", ratio)
	fmt.Printf("EBT is %0.2f, Profit is %0.4f, and Ratio was %0.4f", earningBeforeTax, profit,  ratio)


	os.WriteFile("Metrics.txt", []byte("EBT is: " + fmt.Sprint(earningBeforeTax) + ".\nProfit is: " + fmt.Sprint(profit) + ". \nRatio is: " + fmt.Sprint(ratio)), 0644)
}

func calculator(revenue, expenses float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}


