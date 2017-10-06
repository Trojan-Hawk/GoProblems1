// Student Name: Timothy Cassidy
// Student Number: G00333333

package main

// format package import
import "fmt"
// math random package import
import "math/rand"
// time package import
import "time"

func main() {
	// random number generated from function
	numberPicked := randomNum(1, 100)
	// variables
    counter := 0
    guess := 0
    lastGuess := 0
	
    fmt.Printf("I have picked a number between 1 and 100.\n")
    
    // loop that runs for a max of 100 tries
    for counter < 100 {
        fmt.Println("Can you guess what it is...?")
		
		/*
		// input error handling
		for _, c := range s {
        if unicode.IsDigit(c) {
				
			}
		}
        guess = s
		*/
		
        fmt.Scanf("%d", &guess)
        // Buffer flushing
        fmt.Scanf("%d")
		// counter increment
        counter++
		
		// if the guess is correct
		if guess == numberPicked {
            break
        }// if
        // checks current guess against the last guess
        if lastGuess == guess {
            fmt.Println("Same guess as last time entered! Guess again..")
			// counter decrement
            counter--
        }// if
		// if the guess is too low
        if guess < numberPicked {
            fmt.Println("Too low, Try again..")
        }// if
		// if the guess is too high
        if guess > numberPicked {
            fmt.Println("Too high, Try again..")
        }// if
        
		// storing the current guess
        lastGuess = guess
    }// for
	
	// if the guess matches the random number
	// else print incorrect
    if guess == numberPicked {
        fmt.Printf("Correct!! And it only took %d tries.\n", counter)
    } else {
        fmt.Printf("Incorrect!! The number I picked was %d.\n", numberPicked)
		fmt.Println("Please try again!")
    }// if/else
}// main

// random number generator function
func randomNum (min, max int) int {
    // using system clock as a seed to generate a random number
    rand.Seed(time.Now().Unix())
	// returning the random number
    return rand.Intn(max - min) + min
}// randomNum