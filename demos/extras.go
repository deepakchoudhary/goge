package main

// Documentation and examples for the GoGE extras package.
// This includes methods that let you do default Go things in simpler ways.
// First, import goge/extras.
// And if you want to print stuff, we recommend importing fmt.

import(
  "goge/extras"
  "fmt"
)

func main() {
  // We made 3 functions that make getting input easier.
  // The first is InitInput(). Thit stops the user from seeing anything they type and turns input buffering off.
  // BUG - This works only on Unix-based systems. We are working on this.
  extras.InitInput()
  
  // Then, you have to get input.
  // Use GetInput()
  // For example, this will wait for input and then print it out:
  fmt.Println("Type in a character.")
  input := extras.GetInput()
  fmt.Println("You entered: " + input)
  
  // Finally, use StopInput() to undo the effects of InitInput().
  // BUG - Again, this works only on Unix-based systems.
  extras.StopInput()
  
  // Next, use Times() to print out a given string a given number of times.
  // This will print out CHEESE CHEESE CHEESE 
  fmt.Println(extras.Times("CHEESE ", 3))
  
  // You can use IsWindows() to check for the operating system.
  if (extras.IsWindows()) {
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
  
  // And orint it.
  fmt.Println(flippedStr)
  
  // You can also exit with Exit() and exit with an error with Error().
  // The following code will wait for input. If input is RETURN, it will exit cleanly, otherwise it will return an error.
  extras.InitInput()
  exitInput := extras.GetInput()
  if (exitInput == "\n") {
    extras.StopInput()
    extras.Exit()
  } else {
    extras.StopInput()
    extras.Error("you didn't hit return")
  }
}