package main

// Documentation and examples for the GoGE extras package.
// This includes methods that let you do default Go things in simpler ways.
// First, import goge/extras.
// And if you want to print stuff, we recommend importing fmt.

import (
	"fmt"
	"goge/extras"
)

func main() {
	// To get input...
	// Use GetInput()
	// For example, this will wait for input and then print it out:
	fmt.Println("Type in a character.")
	input := extras.GetInput()
	fmt.Println("You entered: " + input)

	// Next, use Times() to print out a given string a given number of times.
	// This will print out CHEESE CHEESE CHEESE
	fmt.Println(extras.Times("CHEESE ", 3))

	// You can use IsWindows() to check for the operating system.
	if extras.IsWindows() {
		fmt.Println("Windows!")
	} else {
		fmt.Println("Not Windows!")
	}

	// There can be a slight delay.
	// This is in milliseconds.
	extras.Delay(3000)

	// You can clear the screen:
	extras.ClearScreen()

	// And you can use Flip() to flip a string.
	// This will also switch characters like <> and {}.
	// So Flip("{[>>}") will output "{<<]}".
	// Let's try it!
	str := "{[>>}"
	flippedStr := extras.Flip(str)

	// And print it.
	fmt.Println(flippedStr)

	// You can also exit with Exit() and exit with an error with Error().
	// The following code will wait for input. If input is \n (ENTER), it will exit cleanly, otherwise it will return an error.
	exitInput := extras.GetInput()
	if exitInput == "\n" {
		extras.Exit()
	} else {
		extras.Error("you didn't hit return")
	}
}
