package rounding

import (
	"fmt"
	"math"
	"strconv"
)

type FloorRounder struct{}

type CeilRounder struct{}

type MathRounder struct{}

func (f FloorRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}
	pow := math.Pow(10, float64(places))
	return math.Floor(value*pow) / pow
}

func (f CeilRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}
	pow := math.Pow(10, float64(places))
	return math.Ceil(value*pow) / pow
}

func (f MathRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}
	pow := math.Pow(10, float64(places))
	scaled := value * pow
	if scaled >= 0 {
		return math.Floor(scaled+0.5) / pow
	}
	return math.Ceil(scaled-0.5) / pow
}

func FormatFloat(value float64, places int) string {
	if places < 0 {
		places = 0
	}

	format := fmt.Sprintf("%%.%df", places)
	return fmt.Sprintf(format, value)
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func TruncateToInt(value float64) int {
	return int(value)
}
