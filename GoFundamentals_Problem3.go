// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"

// main function
func main() {

	// for loop that runs 100 times
    for i := 1; i <= 100; i++ {
		// if i is divisable by 3 and 5 print "Fizzbuzz"
		// or if i is divisable by 3 print "Fizz"
		// or if i is divisable by 5 print "Buzz"
		// if none print value for i
		if i % 3 == 0 && i % 5 == 0 {
		  fmt.Println("FizzBuzz")
		} else if i % 3 == 0{
		  fmt.Println("Fizz")
		} else if i % 5 == 0{
		  fmt.Println("Buzz")
		} else {
		 fmt.Println(i)
		}
	}// for
	
}// main