// Regular Expressions
// Edward Eldridge (G00337490)

package main

import(
	"fmt" // For taking in user input (https://golang.org/pkg/fmt/)
 	//"bufio" // For input/output (https://golang.org/pkg/bufio/)
 	"math/rand" // For our RNG (https://golang.org/pkg/math/rand/)
	"time" // For our rng (https://golang.org/pkg/time/)
	//"os" // For our reader
)
		// Function that takes a single string as an input
		// and returns a random string from an array
		func ElizaResponse() string{
			// Variables
			var outputArray [3]string
			var userInput string

			// Initlize our array
			outputArray[0]="I’m not sure what you’re trying to say. Could you explain it to me?"
			outputArray[1]="How does that make you feel?"
			outputArray[2]="Why do you say that?"

			// Prompt user for input
			// Create new reader so we can read strings
			fmt.Println("Please enter a sentence: ")
			fmt.Scanf("%s", &userInput)

			// Set a seed for our RNG using time in nano seconds
			rand.Seed(time.Now().UnixNano())

			// Return one of the 3 strings held in our string array, chosen at random
			return outputArray[rand.Intn(3)]
			



			test
		}
 func main() {
	ElizaResponse();
 }

