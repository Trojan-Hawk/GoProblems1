package main

// format package import
import "fmt"
// math/big package import
import "math/big"

// Source
// https://play.golang.org/p/feacvk4P4O
// https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion

// main function
func main() {
	// setting the variable number as a big int with value of 100
	number := big.NewInt(100)
	// calling the factorial function in a print line
	number = (factorial(number))
	// printing the sum
	fmt.Println("The sum of 100! is:", sum(number))
	
}// main()

// function that recursivly calls itself until 0 is reached
func factorial(num *big.Int) (result *big.Int) {
	zero := big.NewInt(0)
	minus1 := big.NewInt(1)
	
	// check to see if the number is less than 0
	if num.Cmp(zero) == -1 {
		result = big.NewInt(1)
	}
	// check to see if the number is equal to 0
	// if it is we set result to 1
	// if not call the function again
	if num.Cmp(zero) == 0 {
		// set result equal to 1
		result = big.NewInt(1)
	} else {
		// set result as a bigInt variable
		result = new(big.Int)
		// Set result equal to the value of num
		result.Set(num)
		// multiply number by the itself - 1
		// then pass it through to the function again
		result.Mul(result, factorial(num.Sub(num, minus1)))
	}
	return
}// factorial()

// 
func sum(number *big.Int) *big.Int{
	// variables
	// bigInt of 10 to get the remainder for the addition
	ten := big.NewInt(10)
	// bigInt variable to store the addition
	sum := big.NewInt(0)
	// used to store the remainder
	mod := big.NewInt(0)
	
	// for loop to parse the big int value
	for ten.Cmp(number)<0 {
		// addition of the remainder of the number / 10 and the sum
		sum.Add(sum, mod.Mod(number,ten))
		// division of the number by 10 
		// to remove the number we have already added to the sum
		number.Div(number,ten)
	}// for
	
	// addition of the final number
	sum.Add(sum,number)
	
	// returning the sum
	return sum
}// sum()











