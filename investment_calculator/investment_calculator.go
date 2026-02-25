package main

import (
	"flag"
	"fmt"
	"math"
)

// calculateMaturityValue computes the compound value of an investment
// after a given number of years at the expected annual return rate.
func calculateMaturityValue(investmentAmount uint, expectedReturnRate float64, years int) float64 {
	if expectedReturnRate < 0 {
		panic("expected annual return rate must be non-negative")
	}

	if years < 0 {
		panic("years must be non-negative")
	}

	return float64(investmentAmount) * math.Pow(1.0 + expectedReturnRate / 100.0, float64(years))
}

// adjustForInflation discounts a future amount back to today's money
// using an annual inflation rate and the same investment horizon.
func adjustForInflation(amount float64, inflationRate float64, years int) float64 {
	if years < 0 {
		panic("years must be non-negative")
	}

	return amount / math.Pow(1.0 + inflationRate / 100.0, float64(years))
}

func main() {
	const (
		defaultInvestmentAmount   uint    = 1000
		defaultExpectedReturnRate float64 = 5.5
		defaultYears              int     = 10
		defaultInflationRate      float64 = 6.5
	)

	investmentAmount := flag.Uint("amount", defaultInvestmentAmount, "initial investment amount")
	expectedReturnRate := flag.Float64("rate", defaultExpectedReturnRate, "expected annual return rate (percent)")
	years := flag.Int("years", defaultYears, "investment duration in years")
	inflationRate := flag.Float64("inflation", defaultInflationRate, "annual inflation rate (percent)")

	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Warning: the following arguments are currently unsupported:")
		for _, arg := range remainingArgs {
			fmt.Printf("  %s\n", arg)
		}
		fmt.Println()
	}

	maturityValue := calculateMaturityValue(*investmentAmount, *expectedReturnRate, *years)
	inflationAdjustedMaturityValue := adjustForInflation(maturityValue, *inflationRate, *years)

	fmt.Println()
	fmt.Println("                     Maturity value =", maturityValue)
	fmt.Println("Maturity value after inflation adj. =", inflationAdjustedMaturityValue)
	fmt.Println()
}
