package main

import (
	"fmt"
	"math"
)

func main() {
	var inverstmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 5

	var futureValue = float64(inverstmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
