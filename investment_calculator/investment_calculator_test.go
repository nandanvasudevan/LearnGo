package main

import (
	"math"
	"testing"
)

func floatEquals(t *testing.T, got, want, tolerance float64) {
	t.Helper()
	if math.Abs(got-want) > tolerance {
		t.Fatalf("got %v, want %v (tolerance %v)", got, want, tolerance)
	}
}

func TestCalculateMaturityValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		investmentAmount  uint
		expectedRate      float64
		years             uint
		want              float64
	}{
		{
			name:             "zero years returns principal",
			investmentAmount: 1000,
			expectedRate:     5.5,
			years:            0,
			want:             1000,
		},
		{
			name:             "zero rate no growth",
			investmentAmount: 1500,
			expectedRate:     0,
			years:            10,
			want:             1500,
		},
		{
			name:             "positive rate and years",
			investmentAmount: 1000,
			expectedRate:     5.5,
			years:            10,
			want:             float64(1000) * math.Pow(1.0+5.5/100.0, 10),
		},
	}

	const tolerance = 1e-6

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := calculateMaturityValue(tc.investmentAmount, tc.expectedRate, tc.years)
			floatEquals(t, got, tc.want, tolerance)
		})
	}
}

func TestAdjustForInflation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		amount        float64
		inflationRate float64
		years         uint
		want          float64
	}{
		{
			name:          "zero years no change",
			amount:        2000,
			inflationRate: 3.2,
			years:         0,
			want:          2000,
		},
		{
			name:          "zero inflation no change over time",
			amount:        5000,
			inflationRate: 0,
			years:         15,
			want:          5000,
		},
		{
			name:          "positive inflation discounts value",
			amount:        10000,
			inflationRate: 2.5,
			years:         5,
			want:          10000 / math.Pow(1.0+2.5/100.0, 5),
		},
	}

	const tolerance = 1e-6

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := adjustForInflation(tc.amount, tc.inflationRate, tc.years)
			floatEquals(t, got, tc.want, tolerance)
		})
	}
}

func TestAdjustForInflationInverseOfMaturityValueWhenRatesMatch(t *testing.T) {
	t.Parallel()

	const (
		investmentAmount  uint    = 2500
		rate              float64 = 4.0
		years             uint    = 12
		tolerance                 = 1e-6
	)

	maturity := calculateMaturityValue(investmentAmount, rate, years)
	got := adjustForInflation(maturity, rate, years)

	floatEquals(t, got, float64(investmentAmount), tolerance)
}

