package main

import (
	"fmt"
	"math"
)

func main() {
	var inverstmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 5

	var futureValue = inverstmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
