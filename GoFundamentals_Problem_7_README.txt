=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================
						
						Task:
	Write a function that tests whether a string is a palindrome. 
	A palindrome is a string that reads the same in reverse, 
					e.g. radar.
	
					  Procedure:
Start by creating a .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.
Start by declaring the package name, in this it will be 'main'.
Then import the 'fmt' and the 'string' package.
The format package alllows us to call format on any print lines we have.
The string package alllows us to use string methods like ToLower.

First we call the main function.
Then declare a string variable.
Next we store a user input string into that variable, then set it 
all to lowercase using the .ToLower method.

Next is an if else statement that prints out palindrome validation to the user.
Inside the if statement the palindrome method is called.
This plndrmCheck method returns true if the input string is a palindrome.
It doest this by using the center of the string as a loop control,
then it checks if the first and last letter are the same and continues
inward until it finishes or finds unmatching characters. 