// Regular Expressions
// Edward Eldridge (G00337490)

package main

import (
	"bufio"
	"fmt"
	"math/rand" // For taking in user input (https://golang.org/pkg/fmt/)
	"regexp"
	"strings"
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
		"I'm not sure what you're trying to say, could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	// Set a variable to father so we can compare the userInput against it
	father := regexp.MustCompile(`(?i)\bfather\b`)
	iAmAMatch := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

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

	// If the userInput is any variation of the word father, enter the statement
	if father.MatchString(userInput) {

		// Print response
		fmt.Println("Why don't you tell me more about your father?")
	} else {

		// If any variation of the words 'I am' match any of the userInput, enter the statement
		if iAmAMatch.MatchString(userInput) {

			// If the userInput contain's 'your', enter the statement
			if strings.Contains(strings.ToLower(userInput), "you're") {

				// Create a variable called yourReplace that will be used to replace 'you're' with me
				yourReplace := regexp.MustCompile(`you're`)

				userInput = yourReplace.ReplaceAllString(userInput, "my")

			}

			// If the userInput contain's 'me', enter the statement
			if strings.Contains(strings.ToLower(userInput), "me") {

				// Create a variable called yourReplace that will be used to replace 'me' with me
				meReplace := regexp.MustCompile(`me`)

				userInput = meReplace.ReplaceAllString(userInput, "you")

			}

			// If the userInput contain's 'you', enter the statement
			if strings.Contains(strings.ToLower(userInput), "you") {

				// Create a variable called yourReplace that will be used to replace 'you' with me
				youReplace := regexp.MustCompile(`you`)

				userInput = youReplace.ReplaceAllString(userInput, "I")

			}

		} else {
			// Replace userInput with "How do you are know you are"
			replacer := iAmAMatch.ReplaceAllString(userInput, "How do you know you are")

			// Print result
			fmt.Println(replacer + "?")

		}
		// Return random answer from outputArray
		fmt.Println("Random Output: " + outputArray[rand.Intn(len(outputArray))])
	}

	// Return the randomString
	return userInput
}

func main() {
	ElizaResponse()
}
