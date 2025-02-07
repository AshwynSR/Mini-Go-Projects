package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueToBeStored := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueToBeStored), 0744)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find value file")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("Failed to parse the data in file")
	}
	return value, nil
}
