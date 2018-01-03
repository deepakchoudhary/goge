package main

// Documentation and examples for the optional GoGE score keeper.
// First, import goge/score.
// And if you want to print stuff, we recommend importing fmt.

import(
  "goge/score"
  "fmt"
)

func main() {
  // Initialize a permanent score keeper.
  // (i.e. for the whole game)
  // This contains 4 things:
  //  Current - The current total score.
  //  Lowest - The lowest score.
  //  Highest - The highest score.
  //  GamesPlayed - The number of games played.
	permanentScore := score.NewPerm()
  
  // Initialize a temporary score keeper.
  // (i.e. for the current round)
  // This only contains 1 thing:
  //  Score: The current score.
	temporaryScore := score.NewTemp()

  // You can add to and subtract from the temporary score.
	temporaryScore.Add(10)
	temporaryScore.Subtract(2)
	
  // You can set it.
	temporaryScore.Set(5)
  
  // And you can set it back to 0.
	temporaryScore.Reset()
  
  // P.S. They're chainable.
  temporaryScore.Set(2).Add(3).Subtract(1)
  
	// You can access the score.
  currentScore := temporaryScore.Score
  
  // And print it.
  fmt.Println(currentScore)
  
  // When the round is done, you can add it to the permanent score keeper.
  // This will also reset the temporary score.
  // And it will automatically calculate the high, low, and current scores for you, as well as adding 1 to the number of games played.
	permanentScore.Add(temporaryScore)
  
  // If you want, you can access those.
  totalScore := permanentScore.Current
  lowScore := permanentScore.Lowest
  highScore := permanentScore.Highest
  gamesPlayed := permanentScore.GamesPlayed
  
  // And print these, too.
  fmt.Println(totalScore)
  fmt.Println(lowScore)
  fmt.Println(highScore)
  fmt.Println(gamesPlayed)
  
  // Finally, you can reset the permanent score keeper.
	permanentScore.Reset()
}