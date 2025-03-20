package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChannels[i] = make(chan bool)
		errorChannels[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[i], errorChannels[i])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <- errorChannels[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChannels[i]:
			fmt.Println("Done!")
		}
	}
}
