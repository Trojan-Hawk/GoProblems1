// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"
// strings package import
import "strings"

func main() {
	// variables
	inputString := ""
	
	// user input
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &inputString)
	
	// setting the string to all lowercase
	inputString = strings.ToLower(inputString)
	
	// boolean controlled if/else statement
	if(plndrmCheck(inputString) == true){
		fmt.Println("The string entered is a Palimdrome!.")
	} else {
		fmt.Println("The string entered is not Palimdrome!.")
	}// if/else
}// main

// function to check if the string is a plaindrome
func plndrmCheck(s string) bool {
	// variables
	// center finds the character in the center by
	// by dividing the length by 2
	center := len(s) / 2
	// end finds the last character by finding the length
	// of the string minus 1
	end := len(s) - 1

	// for loop to parse the string and to compare
	// each characters with its corresponding oppisite
	for i := 0; i < center; i++ {
		
		// if any characters don't match return false
		if s[i] != s[end-i] {
			return false
		}
	}// for
	
	// if they all match return true
	return true
}// plndrmCheck