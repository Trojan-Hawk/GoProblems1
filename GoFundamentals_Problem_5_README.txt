=========================================================================================
	Student Name: Timothy Cassidy		Student Number: G00333333
=========================================================================================

Start by creating the .go file.
Open the command prompt and navigate to the .go file directory.
Open up the .go file in a text editor.

Start by declaring the package name, in this it will be 'main'.
Then in the main we import the 'fmt', 'time' and the 'math/rand' package.
The format package alllows us to call format on any print lines we have.
The math/rand package allows us to generate a random number.
The time package allows us to access the system clock,
the system time will be used as the seed value for our random number
to ensure it is always random.

First we call the main function.
Inside this we call a function to generate a random number and store it.
The random number function uses the system clock as a seed to generate 
a random number by using the rand.Seed and the rand.Intn methods.
Next we declare our variables.

After this is a for loop which scans in the users guess,
increments a counter and checks the users input against the randomly generated number.
It prints out a different message depending on what number the user entered.

After the for loop is an if statement which lets the user know their guess
was correct or if they have used up all their tries, then the program terminates.