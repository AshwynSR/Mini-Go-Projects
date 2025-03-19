package prices

import (
	"example/price-calculator/conversion"
	"example/price-calculator/filemanager"
	"fmt"
)

type TaxIncludedPricesJob struct {
	IOManager filemanager.FileManager `json:"-"`    // to exclude this field in output file
	TaxRate           float64 `json:"tax_rate"`
	Prices            []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadPrices() {
	// fm := filemanager.New("prices.txt", "result.json")

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	finalPrices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices = finalPrices
}

func (job *TaxIncludedPricesJob) Process() {
	// fm := filemanager.New("prices.txt", fmt.Sprintf("result-%.0f.json", job.TaxRate*100))
	job.LoadPrices()
	result := make(map[string]string)
	for _, price := range job.Prices {
		finalPrice := price * (1 + job.TaxRate)
		result[fmt.Sprint(price)] = fmt.Sprintf("%.2f",finalPrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	fmt.Println(result)
}

func New(fm filemanager.FileManager, taxRate float64, prices []float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: fm,
		TaxRate: taxRate,
		Prices:  prices,
	}
}
