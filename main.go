package main

import (
	"math"
)

func solveEquation(a, b, c float64) (*float64, *float64) {
	d := b*b - 4*a*c
	if d < 0 {
		return nil, nil
	}
	if d == 0 {
		x1 := -b / (2 * a)
		return &x1, nil
	}
	sqrt := math.Sqrt(d)
	x1 := (-b - sqrt) / (2 * a)
	x2 := (-b + sqrt) / (2 * a)
	return &x1, &x2
}

func main() {
}
