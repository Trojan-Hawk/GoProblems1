// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// time and format package import
import (
	"fmt"
	"time"
	)

// Time And Date main function
func main() {
	
	// Storing the current time and date into a variable
	// Which can then be split up into seperate parts
	curTimeDate := time.Now()
	
	// Formatted print of the current time
    fmt.Println("Current time is: ",curTimeDate.Hour(),":",curTimeDate.Minute(),":",curTimeDate.Second())
	
	// Formatted print of the current date
	fmt.Println("Current Date is: ",curTimeDate.Day(),curTimeDate.Month(),curTimeDate.Year())
}// main