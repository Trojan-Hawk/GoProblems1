=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================
						
						Task:
	Write a function that merges two sorted lists into a new 
	sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
	
					  Procedure:
Start by creating a .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.
Start by declaring the package name, in this it will be 'main'.
Then import the 'fmt' package.
The format package alllows us to call format on any print lines we have.

First we call the main function.
Then declare two int arrays.
Then we print the two arrays and call the merge function on them.

The merge function merges two sorted arrays together and then returns the new array.
It does this by looping through both arrays until one of them is fully parsed.
Inside the loop are some if statements.
If array1 is empty, append array2 to new array.
If array2 is empty, append array1 to new array.
If array1 element 0 is less than or equal to array2 element 0,
append element 0 to new array and shorten array1.
Otherwise append element 0 of array2 to new array and shorten array2.