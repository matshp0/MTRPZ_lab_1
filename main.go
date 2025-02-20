package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	Italic   = "\033[3m"
	Green    = "\033[32m"
	Combined = Reset + Bold + Italic + Green
)

func validateFloat(input string) (float64, error) {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}
	return val, nil
}

func processInput() (float64, float64, float64) {
	var arr [3]string
	var result [3]float64
	var err error
	params := [3]string{"a", "b", "c"}
	fmt.Println("Enter equation parameters:")

	for i, el := range params {
		for {
			fmt.Printf(Reset+"%v = "+Combined, el)
			fmt.Scan(&arr[i])
			result[i], err = validateFloat(arr[i])
			if err != nil {
				fmt.Printf(Reset+"Expected a valid real number, got '%v' instead\n", arr[i])
				continue
			}
			break
		}
	}
	fmt.Print(Reset)
	return result[0], result[1], result[2]
}

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

func logOutput(x1, x2 *float64) {
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

func main() {
	var x1, x2 *float64
	fmt.Printf("\033[0m")
	a, b, c := processInput()
	x1, x2 = solveEquation(a, b, c)
	logOutput(x1, x2)
}
