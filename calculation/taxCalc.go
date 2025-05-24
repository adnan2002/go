package calculation

type TaxResult struct {
	BasePrice int                `json:"base_price"`
	AfterTax  map[string]float64 `json:"after_tax"`
}

func GetTaxResult(numbers []int) []TaxResult {
	taxeRates := [3]float64{0.0, 10.0, 20.0}

	results := make([]TaxResult, len(numbers)) // No need for cap, len is enough

	for i, value := range numbers {
		results[i].BasePrice = value // Set BasePrice once per TaxResult
		results[i].AfterTax = make(map[string]float64) // Initialize the map here

		for _, val := range taxeRates {

			switch val {
			case 0.0:
				results[i].AfterTax["0%"] = calc(value,val)
			case 10.0:
				results[i].AfterTax["10%"] = calc(value,val)
			case 20.0:
				results[i].AfterTax["20%"] = calc(value,val)
			}
		}
	}

	return results
}

func calc(basePrice int, taxRate float64) float64 {
	return float64(basePrice) * (1 + taxRate/100)
}