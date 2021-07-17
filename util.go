package main

import "math"

// helps compare the float values
// we need our values relatively equal
// we can consider our values equal
// assuming some rounding error
func relativelyEqual(x0, x1 float64) bool {
	const maxRelativeDiff = 0.09
	return math.Abs(x0-x1) <= maxRelativeDiff
}

// round a value to 2 decimal places
func toTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100)/100
}
