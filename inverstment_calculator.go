package main

import (
	"fmt"
	"math"
)

func main() {
	const infalationRate = 2.1
	var inverstmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("inverstment Amount : ")
	fmt.Scan(&inverstmentAmount)

	fmt.Print("Enter Years  : ")
	fmt.Scan(&years)

	fmt.Print("Enter Return Rate  : ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = inverstmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+infalationRate/100, years)
	//formated string
	fmt.Println("Future Value :", futureValue)
	fmt.Println("Future Real Value :", futureRealValue)

}
