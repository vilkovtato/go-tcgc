package main

import (
	"fmt"

	"github.com/vilkovtato/projectcalculator/cmdmanager"
	"github.com/vilkovtato/projectcalculator/filemanager"
	"github.com/vilkovtato/projectcalculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		iomanager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(iomanager, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println((err))
		}
	}

	for _, taxRate := range taxRates {
		iomanager := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(iomanager, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println((err))
		}
	}
}