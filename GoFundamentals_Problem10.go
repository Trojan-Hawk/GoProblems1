// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"

func main() { 
		// variables
        input := ""
		newString := ""
		
        // when the user inputs a space it's counted as an end of input
        fmt.Println("Please input a single word (No spaces!)")
        fmt.Scanf("%s", &input)
        // Rune turns the string letters into unicode 
        // Getting the unicode code points.
		
		// storing the length of the entered string in a variable
		r := len(input)
		
		// for loop controlled by the length of the string
		// takes the last letter of the string and stores it in a new string
		// keeps moving left along the string and storing it's values
		for i := 1; i <= r; i++{ 
			// appending the characters onto the new string variable
			newString += string([]rune(input)[r - i])
        }// for 
		
		// Print of the new reversed string
		fmt.Printf("\n%s reversed is: %s", input, newString)

}// main