// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"
// math package import
import "math"

func main() {
	// constants
	number := 20.0
	
	// variables	
	var guess float64
	
	// loop to aproximate the square root using Newton’s method
	// we start this at 1
	// if the guess is equal to the last the loop ends
	for guess = 1.0; guess != squareRoot(guess, number); guess = squareRoot(guess, number) {
		fmt.Printf("Current square root approximation: %12.8f\n", guess)
	}// for

	// Print of the square root approximation using Newton’s method
	fmt.Printf("\nThe approximate square root of %f is %f \n", number, guess)
	
	// Print of the exact square root of number
	fmt.Printf("\nThe actual square root of %f is %f", number, math.Sqrt(number));
}// main

// function with the square root equation
func squareRoot(guess float64, number float64) float64{
	// equation of Newton's method (((x * x) - x) / (2 * x))
	return guess - (((guess * guess) - number) / (2 * guess))
}// squareRoot