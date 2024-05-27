package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount float64
	// var years float64
  // expectedReturnRate := 5.5
	
	investmentAmount := outputText("Investment Amount: ")
	years := outputText("Years: ")
	expectedReturnRate := outputText("Expected Retun Rate: ")

	// fmt.Print("Investment Amount: ")
	// outputText("Investment Amount: ")
	// fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	// outputText("Expected Return Rate: ")
	// fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	// outputText("Years: ")
	// fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	formattedFV := outputTextFormat("Future Value: %.2f\n", futureValue)
	formattedRFV := outputTextFormat("Future Real Value: %0.2f\n", futureRealValue)

	// fmt.Printf(`Future Value: $%.2f
	//
	// 	Future Real Value: $%.2f`, futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) (input float64) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func outputTextFormat(text string, variable float64) (string) {
	return fmt.Sprintf(text, variable)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}


