package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strArray []string) ([]float64, error) {
	prices := []float64{}
	for _, price := range strArray {
		val, err := strconv.ParseFloat(price, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}
		prices = append(prices, val)
	}

	return prices, nil
}
