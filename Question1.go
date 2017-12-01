// Regular Expressions
// Edward Eldridge (G00337490)

package main

import (
	"bufio"
	"fmt"
	"math/rand" // For taking in user input (https://golang.org/pkg/fmt/)
	// For our RNG (https://golang.org/pkg/math/rand/)
	"os"
	"time" // For our rng (https://golang.org/pkg/time/)
)

// Function to generate random number within a range
// Adapted from (http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html)
//func random(min, max int) int {

// Set a seed for our RNG using time in nano seconds

// Return our random integer to our main function
// where we can set it's parameters
//return rand.Intn(max-min) + min
//}

// Function that takes a single string as an input
// and returns a random string from an array of strings
func ElizaResponse() string {

	rand.Seed(time.Now().Unix())

	// Initlize an array of strings that will be randomly chosen
	outputArray := []string{
		"I'm npot sure what you're trying to say, could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//Put the random string in our array into a variable called randomString
	//outputArray[randomNumber] = randomString

	// Prompt user for input
	// Create new reader so we can read strings
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// Stop reading when a newLine (\n) is read (Windows problem)
	userInput, _ := reader.ReadString('\n')

	// Test user input
	fmt.Println()
	fmt.Println("User Input:" + userInput)

	// Return random answer from outputArray
	fmt.Println("Random Output: " + outputArray[rand.Intn(len(outputArray))])
	// Return the randomString
	return userInput
}

func main() {
	ElizaResponse()
}
