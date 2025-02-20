# Quadratic Equation Solver

## Description
This program solves quadratic equations of the form:

```
ax^2 + bx + c = 0
```
It supports two modes of input:
1. **Interactive Mode**: The user enters the coefficients (`a`, `b`, and `c`) manually.
2. **File Mode**: The program reads the coefficients from a file.

The program validates user input and computes the roots of the equation using the quadratic formula.

## How to Run

### 1. Interactive Mode
Run the program without arguments to enter the coefficients manually:
```sh
./program
```
The program will prompt you to input the values for `a`, `b`, and `c`.

### 2. File Mode
Run the program with a file containing the coefficients:
```sh
./program input.txt
```
The `input.txt` file should contain three space-separated numbers, e.g.,:
```
1 -3 2
```

## Output
The program prints the roots of the quadratic equation:
- If there are two real roots: `Found two roots: x1, x2`
- If there is one real root: `Found one root: x`
- If there are no real roots: `No roots found`

## Dependencies
- The program is written in Go and requires a Go runtime.
- Ensure Go is installed on your system.

## Compilation
To compile the program, run:
```sh
go build -o program main.go
```

## Error Handling
- If the user enters an invalid number, they will be prompted to retry.
- If the file contains invalid data, the program will terminate with an error.


