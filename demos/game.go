package main

// An example of how to make a simple game with GoGE.
// We will use all of goge's functions.
// Import goge, goge/extras, and goge/score first.
// Also import fmt for printing information to the terminal.
// NOTE - Due to a bug, input will only work on Unix-based systems. This is being fixed.

import (
	"fmt"
	"goge"
	"goge/extras"
	"goge/score"
)

// Create an exit function.
// This will ask users if they want to continue or quit.
func exit() {
	// Ask them for input.
	fmt.Println("Hit ENTR to continu or X to qut.")

	// Collect the input.
	input := ""
	for {
		input = extras.GetInput()
		if input == "\n" {
			// Continue.
			return
		} else if input == "x" {
			// Stop the input, clear the screen, and then exit.
			extras.ClearScreen()
			extras.Exit()
		}
	}
}

// Create a permanent and temporary score.
var perm = score.NewPerm()
var temp = score.NewTemp()

func main() {
	// Create a 5x5 grid.
	// Each model is 4 characters long.
	grid := goge.Init(4, 5)

	// Create a framer with a 10 millisecond delay between each frame.
	// We use this low delay so the gae seems fluid.
	framer := grid.Framer(10)

	// Create the creatures.
	creature := goge.New("You", "_O-*")
	cheese := goge.New("Cheas", " [> ")

	// Put them in the grid at 1, 1 and 5, 5.
	grid.Use(creature, 1, 1)
	grid.Use(cheese, 5, 5)

	// Clear the screen and print game info.
	extras.ClearScreen()
	fmt.Println("The gole of the gaim is to eat the Peic of Cheas.")
	fmt.Println("Us W, A, S, and D to cuntrol yur charcter.")

	// Print the scoring info.
	fmt.Print("\nHi scor: ")
	fmt.Print(perm.Highest)
	fmt.Print(".\nLow scor: ")
	fmt.Print(perm.Lowest)
	fmt.Print(".\nCurent totle scor: ")
	fmt.Print(perm.Current)
	fmt.Print(".\nYuve one ")
	fmt.Print(perm.GamesPlayed)
	fmt.Println(" gaims\n")

	// Check with the user if they want to continue.
	exit()

	// Display the blank grid.
	framer.Frame()

	// Give the user a starting score of 100.
	temp.Set(100)

	// And print it out.
	fmt.Print("Current score: ")
	fmt.Print(temp.Score)

	// Main game loop:
	input := ""
	for {
		// Get input.
		input = extras.GetInput()
		if input == "w" {
			// W key pressed: Move up.
			// (You can access a creature's current position with creature.Row and creature.Column.)
			grid.Up(creature.Row, creature.Column)
		} else if input == "a" {
			// A key pressed: Move to the left.
			grid.Left(creature.Row, creature.Column)
		} else if input == "s" {
			// S key pressed: Move down.
			grid.Down(creature.Row, creature.Column)
		} else if input == "d" {
			// D key pressed: Move to the right.
			grid.Right(creature.Row, creature.Column)
		} else if input == "M" {
			// SHIFT+M keys pressed: "Cheat" and gain 100000 points.
			temp.Add(100000)
		} else {
			// Otherwise, continue with the loop.
			continue
		}

		// Subtract 10 from the user's score.
		temp.Subtract(10)

		if grid.Has("Cheas") {
			// If the cheese is in the grid...
			if temp.Score == 0 {
				// If the score is zero, you lost.
				extras.ClearScreen()
				fmt.Println("You lost.")

				// Check with the user if they want to continue.
				exit()

				// Play the game again.
				main()
			} else {
				// Otherwise, mark a frame.
				framer.Frame()

				// And show the user their score.
				fmt.Print("Current score: ")
				fmt.Print(temp.Score)
			}
		} else {
			// Otherwise, the user won!
			break
		}
	}

	// Add their score to the permanent score.
	extras.ClearScreen()
	perm.Add(temp)

	// Tell them that they won.
	fmt.Println("You one! You aet the Cheas.")

	// Check with the user if they want to continue.
	exit()

	// Play the game again.
	main()
}
