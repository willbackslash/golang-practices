package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumentFunction(a, b int, c int) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPalindrome(text string) bool {
	var invertedString string
	for i := len(text) - 1; i >= 0; i-- {
		invertedString += strings.ToLower(string(text[i]))
	}

	return invertedString == strings.ToLower(text)
}

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

	// fmt
	helloMessage := "Hello"
	worldMessage := "World"
	number := 500

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	fmt.Printf("%s %d\n", helloMessage, number)
	fmt.Printf("%v %v\n", helloMessage, number)

	// Sprintf
	message := fmt.Sprintf("%s %d", helloMessage, number)
	fmt.Println(message)

	// data type print
	fmt.Printf("type helloMessage: %T\n", helloMessage)
	fmt.Printf("type number: %T\n", number)

	// functions and anonymous functions
	normalFunction("Hello World")
	tripleArgumentFunction(1, 2, 3)
	value := returnValue(2)
	fmt.Println("returnValue:", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("doubleReturn:", value1, value2)
	valuea, _ := doubleReturn(2)
	fmt.Println("doubleReturn:", valuea)

	// conditional for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for ever
	//counterForever := 0
	//for {
	//	fmt.Println("Forever", counterForever)
	//	counterForever++
	//}

	// if
	valor1 := 1
	valor2 := 2

	if valor1 == valor2 {
		fmt.Println("valor1 == valor2")
	} else {
		fmt.Println("valor1 != valor2")
	}

	// and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// or
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// cast string to int
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	// switch
	module := 5 % 2
	switch module {
	case 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}

	switch module := 5 % 2; module {
	case 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}

	// switch without condition
	switchValue := 200
	switch {
	case switchValue > 100:
		fmt.Println("value > 100")
	case switchValue < 0:
		fmt.Println("value < 0")
	default:
		fmt.Println("no condition")
	}

	// defer
	defer fmt.Println("Hello")
	fmt.Println("World")

	// continue, break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("is 2")
			continue
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 8 {
			fmt.Println("is 8")
			break
		}
	}

	// arrays and slices
	// arrays
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	// slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))
	// slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// append element
	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	// append slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	// iterate over slice
	myslice := []string{"hello", "world", "how", "are", "you?"}
	for i, value := range myslice {
		fmt.Println(i, value)
	}

	// using just the index
	myslice2 := []string{"hello", "world", "how", "are", "you?"}
	for i := range myslice2 {
		fmt.Println(i)
	}

	fmt.Println("isPalindrome: racecar", isPalindrome("racecar"))
	fmt.Println("isPalindrome: coke", isPalindrome("coke"))
	fmt.Println("isPalindrome: Ama", isPalindrome("Ama"))

	// maps
	m := make(map[string]int)
	m["Jose"] = 14
	m["Juan"] = 15
	fmt.Println(m)

	// iterate over a map
	for key, value := range m {
		println(key, value)
	}

	// searc a value in map
	// var ok let us know if the key exists
	foundValue, ok := m["Jose"]
	fmt.Println(foundValue, ok)
}
