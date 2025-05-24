package main
import (
	"example.com/app/input"
	"example.com/app/calculation"
	"example.com/app/saveJson"
)
func main(){
	numbers := input.GetPricesArray("prices.txt")
	taxResult := calculation.GetTaxResult(numbers)

	saveJson.Save(taxResult)

}