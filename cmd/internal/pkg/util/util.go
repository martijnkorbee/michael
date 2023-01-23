package util

import "math"

func RoundFloat(num float64, precision int) float64 {
	t := math.Pow(10, float64(precision))
	return math.Round(num*t) / t
}
