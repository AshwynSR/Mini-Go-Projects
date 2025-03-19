package main

import (
	"example/price-calculator/filemanager"
	"example/price-calculator/prices"
	"fmt"
)

func main() {

	priceList := []float64{100, 200, 500}
	taxRates := []float64{0, 0.05, 0.1, 0.2}
	// result := make(map[float64][]float64)

	for _, tax := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result-%.0f.json", tax*100))
		job := prices.New(fm, tax, priceList)
		job.Process()
	}

	// fmt.Println(result)
}
