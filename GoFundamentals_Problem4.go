package main

// format package import
import "fmt"
import "math/big"

// Hello World main function
func main() {
	
	// Variables
	startInt := 100
	factorial := new(big.Int)
	
	// for loop that runs 100 times
    for i := 1; i <= startInt; i++ {
		factorial = factorial * (startInt - 1)
		fmt.Println(i, factorial)
	}// for
	
}// main()