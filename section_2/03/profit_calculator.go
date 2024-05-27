package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ebtResultsLocalFile = "ebtResult.txt"
const profitResultsLocalFile = "profitResult.txt"
const ratioResultsLocalFile = "ratioResult.txt"


func writeResultToLocalFile(ebt, profit, ratio float64) {
	ebtText := fmt.Sprint(ebt)
	profitText := fmt.Sprint(profit)
	ratioText := fmt.Sprint(ratio)
	os.WriteFile(ebtResultsLocalFile, []byte(ebtText), 0644)
	os.WriteFile(profitResultsLocalFile, []byte(profitText), 0644)
	os.WriteFile(ratioResultsLocalFile, []byte(ratioText), 0644)

}

func readResultsFromLocalFile(file string) (float64, error){
	data, err := os.ReadFile(file)
	if err != nil {
		return 1, errors.New("Failed to find file")		
	}
	resultText := string(data)
	result, err := strconv.ParseFloat(resultText, 64)
	if err != nil {
		return 1, errors.New("Faild to parce store for vaild data type")
	}
	return result, nil
}

func main() {
	var revenue float64
	var	expenses float64 
	var taxRate float64 
	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax: ")
	ebt, profit, ratio := calulateFinancials(revenue, expenses, taxRate)
	fmt.Printf("\nCURRENT:\n EBT: %.2f\n Profit: %0.2f\n Ratio: %f\n", ebt, profit, ratio)
	}

func calulateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	writeResultToLocalFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func getUserInput(text string) float64 {
	var userInput float64
	for userInput <= 0 {
		fmt.Print(text)
		fmt.Scan(&userInput)
	}
	return userInput
}
