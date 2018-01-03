package main

// An example of how to make a little animation with GoGE.
// It will be a little character chasing around a piece of pizza.
// We'll call him a dragon.
// NOTE - All function involving the grid and the framer are chainable, as you will see. Also, when a character changes direction it's model wil be flipped.
// Import goge first.
// And import goge/extras to clear the screen later.

import(
  "goge"
  "goge/extras"
)

func main() {
  // Initialize a new grid, where each character's ASCII "image" (or model) is 5 characters long and the grid is 3x3.
  grid := goge.Init(5, 3)
  
  // Initialize a framer for the grid.
  // This will actually draw the grid.
  // There will be a 500 millisecond delay between each frame.
	framer := grid.Framer(500)
  
  // Create a new character (or creature) that is named "Dragon".
  // It's model is _O--*.
	dragon := goge.New("Dragon", "_O--*")
  
  // Create a new creature called "Pizza" whose model is [>>.
	pizza := goge.New("Pizza", " [>> ")
  
  // Put the dragon in row 1, column 1.
  // And put the pizzq in row 2, column 3.
	grid.Use(dragon, 1, 1).Use(pizza, 2, 3)
  
  // Mark a frame.
	framer.Frame()
  
  // Move whatever is in row 1, column 1 (the dragon) down.
	grid.Down(1, 1)
  
  // Mark a frame.
	framer.Frame()
  
  // Move whatever is in row 1, column 1 (the dragon) down.
	grid.Right(2, 1)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Down(2, 3).Right(2, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Left(3, 3).Down(2, 3)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Left(3, 2).Left(3, 3)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Up(3, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Right(3, 1)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Down(2, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Change the model of whatever is at 3, 2 (the dragon) to _OO-*.
  // (Mutate it.)
	grid.Mutate("_OO-*", 3, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Mutate the dragon's model to _OOO*.
  // Now it's really fat!
	grid.Mutate("_OOO*", 3, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Up(3, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Up(2, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Move some thing(s).
	grid.Left(1, 2)
  
  // Mark a frame.
	framer.Frame()
  
  // Remove the dragon from the view.
	grid.Empty(1, 1)
  
  // Mark a frame.
	framer.Frame()
  
  // Clear the screen.
  extras.ClearScreen()
}