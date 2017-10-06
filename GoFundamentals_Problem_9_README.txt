=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================
						
							Task:
	Write a function to calculate the square root of a number using Newton’s 
	method. Newton’s method is to approximate the square root of a number x 
	by picking a starting point z and then repeating the following operation.
					z_next = z - ((z*z - x) / (2 * z))
	This is repeated while the values of z_next and z are different. After 
		each iteration the value of z is replaced with that of z_next.

	Run a few tests to determine how close you are to the math.Sqrt value?
	
						  Procedure:
Start by creating a .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.
Start by declaring the package name, in this it will be 'main'.
Then import the 'fmt' and the 'math' package.
The format package alllows us to call format on any print lines we have.

First we call the main function.
Then declare a variable with a value of 20 and an empty guess variable.
Then we run a for loop which calls a square root function,
the loop ends when the function returns the same value.

The squareRoot function uses Newton's method of approximating a square root.
Equation: (((x * x) - x) / (2 * x))
 
After the loop the approximate square root and the actual square root are output.