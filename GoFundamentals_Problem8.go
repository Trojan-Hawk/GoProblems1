// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

import "fmt"

// Source: https://play.golang.org/p/Ma2GXvj3XP

func main() {
	// arrays
	array1 := []int{
		1,7,13,16,21,23,24,27,35,37,54,57,63,74,97,98,
	}
	array2 := []int{
		11,12,31,32,36,42,45,47,53,55,57,61,71,72,73,75,79,89,90,99,
	}
	
	// printing the arrays
	fmt.Printf("Array1:\t%v\nArray2:\t%v\n", array1, array2)
	// calling the merge function and printing
	fmt.Printf("%v\n", Merge(array1, array2))
}// main

// function to merge two arrays
func Merge(array1, array2 []int) []int {
	// new array with length of array1 + array2
	newArray := make([]int, 0, len(array1)+len(array2))

	// for loop that runs until either array1 or array2 is fully parsed
	for len(array1) > 0 || len(array2) > 0 {
		// if array1 is empty, append array2 to new array
		if len(array1) == 0 {
			return append(newArray, array2...)
		}// if
		// if array2 is empty, append array1 to new array
		if len(array2) == 0 {
			return append(newArray, array1...)
		}// if
		
		// if array1 element 0 is less than or equal to array2 element 0,
		// append element 0 to new array and shorten array1
		// else append element 0 of array2 to new array and shorten array2
		if array1[0] <= array2[0] {
			newArray = append(newArray, array1[0])
			array1 = array1[1:]
		} else {
			newArray = append(newArray, array2[0])
			array2 = array2[1:]
		}// if/else
	}// for
	
	// return new array
	return newArray
}// Merge