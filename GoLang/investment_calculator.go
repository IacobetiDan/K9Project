package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Ivestment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println((futureRealValue))

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	//comment format test with printf
	// \n new line
	//for ratio we load the format and the we dispaly
	formatRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)

	fmt.Printf("EBT: %v\n", ebt)
	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Println("Profit:", profit)
	//print ratio
	fmt.Print(formatRatio)

	//for multiple line we us ` ` with this we can use multiple line.
}
