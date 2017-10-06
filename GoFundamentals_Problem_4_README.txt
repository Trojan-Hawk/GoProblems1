=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================
						
							Task:
	Write a function that calculates the sum of the digits of the factorial
	of a number. n! means n x (n − 1) ... x 3 x 2 x 1. For example, 
	10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in 
	the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. Then find the sum of 
					the digits in the number 100!.
	
						  Procedure:
Start by creating the .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.

Start by declaring the package name, in this it will be 'main'.
Then I created two functions, one to calculate the factorial of 100,
and the other to sum each digit of the factorial output.
Then above the main we import the 'fmt' package and the 'math/big' package.
The format package alllows us to call format on any print lines we have.
The math/big package allows us to use the big int method,
this lets us store a value past the max value of an int.
First we call the main function.
Then we declare a bigInt variable number and set it to 100.
We then call the factorial function which calculates the factorial of a number.

The factorial method takes in a bigInt variable and returns a bigInt.
This method calls itself recursivly until the control variable reaches 1.
First we check to see if the number is less than 0, if it is we set result to 1,
which will end the function during the next iteration.
Next we check to see if the number is equal to 0, if it is we set the result to 1,
which ends the function.
If none of these are accessed an else statement is ran, 
this else multiplies the number by itself -1 and calls the function again.
When the function ends it returns the factorial.

Next in the main we call the sum function in a formatted println statement.
Inside this sum function a sum variable is calculated in a for loop,
this for loop takes in the bigInt number, gets the modulus when divided by 10,
then this is added to the sum variable.
Next the number is divided by 10, which removes the the number at the end, 
which we have already added to the sum.
After the loop we add the final digit to the sum and then return sum.