// Created by: Charlotte Jhu
// Created on: April 2023
//
// This program calculates degrees Fahrenheit to degrees Celsius

package main

import "fmt"

func main() {
	// This function calculates degrees Fahrenheit to degrees Celsius
	// variables
	var fahrenheit float64

	// input
	fmt.Println("This program calculates degrees Fahrenheit to degrees Celsius")
	fmt.Println()
	fmt.Print("Enter the temperature in degrees Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	fmt.Println()

	// process
	celsius := (fahrenheit - 32) * 5 / 9

	// output
	fmt.Println("The temperature in Celsius is", celsius)
	fmt.Println("\nDone.")
}
