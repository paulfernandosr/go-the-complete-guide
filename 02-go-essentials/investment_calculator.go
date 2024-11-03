package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)
	//outputs information
	//fmt.Println("Future value:", futureValue)
	//fmt.Printf(`Future Value: %.1f\n

	//Future value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
	//fmt.Println("Future value (adjusted for inflation):", futureRealValue)
	fmt.Print(formatedFV, formatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
