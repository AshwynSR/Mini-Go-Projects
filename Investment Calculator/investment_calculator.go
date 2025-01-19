package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const inflationRate = 5.0
	// const pi float64 = 3.14159
	var investmentAmount float64
	var interestRate = 5.5
	name := "Ashwin"
	var year float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scanln(&investmentAmount)

	fmt.Print("Enter Time Period in Years: ")
	fmt.Scanln(&year)

	fmt.Print("Enter Interest Rate: ")
	fmt.Scanln(&interestRate)

	returns := investmentAmount * math.Pow((1+interestRate/100), year)
	realReturns := returns / math.Pow(1 + inflationRate/100, year)

	fmt.Println(strconv.FormatFloat(returns, 'f', -1, 64) + " for " + name);
	fmt.Println(strconv.FormatFloat(realReturns, 'f', -1, 64) + " for " + name)
}
