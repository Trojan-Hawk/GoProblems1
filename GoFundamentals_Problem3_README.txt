=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================
						
							Task:
	Write a program that prints the numbers from 1 to 100, each on a new line,
	to the console, except for the following conditions. For multiples of 
	three print Fizz instead of the number, and for multiples of five print 
	Buzz. For numbers that are multiples of both three and five print FizzBuzz.
	
						  Procedure:
Start by creating a .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.
Start by declaring the package name, in this it will be 'main'.
Then import the 'fmt' package.
The format package alllows us to call format on any print lines we have.

First we call the main function.
We then run call a for loop that runs 100 times.
Inside the loop is an if statement that prints out something to the console
depending on the value of i.
If i is divisable by 3 and 5, 'FizzBuzz' is output.
If i is divisable by 3, 'Fizz' is output.
If i is divisable by 5, 'Buzz' is output.
And if i is none of the above the value for i is output.