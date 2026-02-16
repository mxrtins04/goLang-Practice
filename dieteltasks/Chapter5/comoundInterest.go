package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0

	for rate := 5.0; rate <= 10.0; rate++ {
		fmt.Printf("\nInterest rate: %.0f%%\n", rate)

		for year := 1; year <= 10; year++ {
			amount := principal * math.Pow(1+rate/100, float64(year))
			fmt.Printf("Year %d: %.2f\n", year, amount)
		}
	}
}

