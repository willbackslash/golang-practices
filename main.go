package main

import "fmt"

func main() {
	// constants
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi:", pi)
	fmt.Println("pi2", pi2)

	// variables int
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// square area
	const squareBase = 11
	squareArea := squareBase * squareBase
	fmt.Println("Area cuadrado:", squareArea)

	// math operators
	// sum
	x := 10
	y := 50
	result := x + y
	fmt.Println("x + y:", result)

	// subtraction
	result = y - x
	fmt.Println("y - x:", result)

	// multiplication
	result = x * y
	fmt.Println("x * y:", result)

	// division
	result = y / x
	fmt.Println("y / x:", result)

	// modulus
	result = y % x
	fmt.Println("y % x:", result)

	// increment
	x++
	fmt.Println("x++:", x)

	// decrement
	x--
	fmt.Println("x--:", x)
}
