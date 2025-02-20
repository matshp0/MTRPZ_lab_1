package main

import (
	"fmt"
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
	a, b, c := 4.0, -5.0, 12.0
	x1, x2 := solveEquation(a, b, c)
	if x1 == nil {
		fmt.Println("No roots found")
		return
	}
	if x2 == nil {
		fmt.Printf("Found one root: %f", *x1)
		return
	}
	fmt.Printf("Found two roots: %f, %f", *x1, *x2)
}
