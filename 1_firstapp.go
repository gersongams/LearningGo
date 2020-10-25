package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// To run this program
	// go fmt 1_firstapp.go // To format it
	// go build 1_firstapp.go // To compile it
	// ./1_firstapp.go to run

	// Otherwise use go run to run the program without generating a an executable file

	fmt.Println("Hello, playground")
	fmt.Println(math.Floor(2.312))
	fmt.Println(strings.Title("my name is gerson"))
	fmt.Println(reflect.TypeOf("GAAA"))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.25412))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(true))

	var isOld bool
	var length, width, height float64
	var customerName string

	isOld = true
	length = 1.2
	width = 5.6
	height = 7.1
	customerName = "Gerson"

	fmt.Println(isOld)
	fmt.Println(length)
	fmt.Println(width)
	fmt.Println(height)
	fmt.Println(customerName)

	var age int = 24
	fmt.Println("My age is", age)

	// Type can be omitted
	var fullName = "Gerson Garrido"
	fmt.Println(fullName)

	// short declaration
	countryName := "France"
	fmt.Println("I would love to live in", countryName)

}

