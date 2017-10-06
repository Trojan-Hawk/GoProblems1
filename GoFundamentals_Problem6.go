// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"

func main() {
	// variables
	largest := 0
	smallest := 999
	// int array
	array := []int{
		20,35,99,66,15,36,22,65,57,79,63,71,19,89,12,41,
	}

	// Array print
	fmt.Printf("\n\tElements in the array:\n%v \n\n", array)

	// Loop to find the largest and smallest array element
  	for _, element := range array {
		// if the current element is bigger than the previous biggest
		if element > largest {
			// set the largest equal to the current element
    		largest = element
    	}// if
		// if the current element is smaller than the previous smallest
		if element < smallest {
			// set the smallest equal to the current element
			smallest = element
		}// if
  	}// for
	
	// largest output
	fmt.Println("The largest array element is:", largest)
	
	// smallest output
	fmt.Println("The smallest array element is:", smallest)
}// main