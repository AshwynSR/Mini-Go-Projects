package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const inflationRate = 5.0
	// const pi float64 = 3.14159
	var investmentAmount = 10000.0
	var interestRate = 9.5
	name := "Ashwin"
	var year = 10.0

	returns := investmentAmount * math.Pow((1+interestRate/100), year)
	realReturns := returns / math.Pow(1+inflationRate/100, year)

	fmt.Println(strconv.FormatFloat(returns, 'f', -1, 64) + " for " + name)
	fmt.Println(strconv.FormatFloat(realReturns, 'f', -1, 64) + " for " + name)
}
