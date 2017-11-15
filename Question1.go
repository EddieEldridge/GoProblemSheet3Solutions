// Regular Expressions
// Edward Eldridge (G00337490)

package main

import(
	"fmt" // For taking in user input (https://golang.org/pkg/fmt/)
 	"bufio" // For input/output (https://golang.org/pkg/bufio/)
 	"rand" // For our RNG (https://golang.org/pkg/math/rand/)
)
	

 func main()
 {
		// Function that takes a single string as an input
		// and returns a random string from an array
		func ElizaResponse() string
		{
			// Variables
			var userInput string; 
			var outputArray [3]string;
			var randomNumber int;

			// Initlize our array
			outputArray[0]="I’m not sure what you’re trying to say. Could you explain it to me?";
			outputArray[1]="How does that make you feel?";
			outputArray[2]="Why do you say that?";

			// Prompt user for input
			fmt.println("Please enter a sentence: ");

			// Create new reader so we can read strings
			reader := bufio.NewReader(os.Stdin) 
			userInput, _ := reader.ReadString('\n')

			// Set a seed for our RNG using time in nano seconds
			randomNumber = rand.Seed(time.Now().UnixNano());

			// Return one of the 3 strings held in our string array, chosen at random
			fmt.Println("Eliza says:", answers[randomNumber(len(outputArray))])

			
		}
 }

