package goge

import (
	"fmt"
	"goge/extras"
)

type __Grid__ struct {
	Grid        [][]*__Creature__
	Chars       int
	Size        int
	__Opacity__ int
}

func (grid *__Grid__) __UpdatePos__() *__Grid__ {
	for row := range grid.Grid {
		for col := range grid.Grid {
			grid.Grid[row][col].Row = row + 1
			grid.Grid[row][col].Column = col + 1
		}
	}

	return grid
}
func (grid *__Grid__) __CalcBorder__() string {
	size := grid.Size
	chars := grid.Chars + 2
	returnString := " |"
	if grid.__Opacity__ == 1 {
		returnString = " |"
	} else {
		returnString = "  "
	}
	for i := 0; i < size; i += 1 {
		if grid.__Opacity__ == 1 {
			returnString += extras.Times("-", (chars))
		} else {
			returnString += extras.Times(" ", (chars))
		}
		if grid.__Opacity__ == 1 {
			returnString += "|"
		} else {
			returnString += " "
		}
	}

	return returnString
}
func __CheckSize__(size int) {
	if size <= 20 {
		return
	} else {
		extras.Error("grid size must be 20 or less.")
	}
}
func (grid *__Grid__) Draw() *__Grid__ {
	fmt.Println(grid.__CalcBorder__())
	for row := range grid.Grid {
		for col := range grid.Grid {
			if grid.__Opacity__ == 1 {
				fmt.Print(" | ")
			} else {
				fmt.Print("   ")
			}
			fmt.Print(grid.Grid[row][col].Model)
		}
		if grid.__Opacity__ == 1 {
			fmt.Println(" |")
		} else {
			fmt.Println("  ")
		}
		fmt.Println(grid.__CalcBorder__())
	}
	__CheckSize__(grid.Size)
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Empty(row int, col int) *__Grid__ {
	grid.__UpdatePos__().Grid[row-1][col-1] = New("__Blank__", extras.Times(" ", grid.Chars), false)
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Has(name string) bool {
	for row := range grid.Grid {
		for col := range grid.Grid {
			if grid.Grid[row][col].Name == name {
				return true
			}
		}
	}
	return false
}
func (grid *__Grid__) Use(creature *__Creature__, row int, col int) *__Grid__ {
	if row > grid.Size {
		extras.Error("creatures cannot be placed outside of the grid")
	}
	if col > grid.Size {
		extras.Error("creatures cannot be placed outside of the grid")
	}
	for row2 := range grid.Grid {
		for col2 := range grid.Grid {
			if row2 == row-1 {
				if col2 == col-1 {
					grid.Grid[row2][col2] = creature
				}
			}
		}
	}
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Mutate(model string, row int, col int) *__Grid__ {
	for row2 := range grid.Grid {
		for col2 := range grid.Grid {
			if row2 == row-1 {
				if col2 == col-1 {
					grid.Grid[row2][col2].Model = model
					dir := grid.Grid[row2][col2].__Direction__
					grid.Grid[row2][col2].__Direction__ = "R"
					grid.Grid[row2][col2].__UpdateDirection__(dir)
				}
			}
		}
	}

	return grid
}
func (grid *__Grid__) Clear() *__Grid__ {
	extras.ClearScreen()

	return grid
}
func (grid *__Grid__) Delay(millis int) *__Grid__ {
	extras.Delay(millis)

	return grid
}
func (grid *__Grid__) Down(row int, col int) *__Grid__ {
	if row != grid.Size {
		grid.Grid[row][col-1] = grid.Grid[row-1][col-1]
		grid.Grid[row-1][col-1] = New("__Blank__", extras.Times(" ", grid.Chars), false)
	}
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Up(row int, col int) *__Grid__ {
	if row != 1 {
		grid.Grid[row-2][col-1] = grid.Grid[row-1][col-1]
		grid.Grid[row-1][col-1] = New("__Blank__", extras.Times(" ", grid.Chars), false)
	}
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Left(row int, col int) *__Grid__ {
	grid.Grid[row-1][col-1].__UpdateDirection__("L")

	if col != 1 {
		grid.Grid[row-1][col-2] = grid.Grid[row-1][col-1]
		grid.Grid[row-1][col-1] = New("__Blank__", extras.Times(" ", grid.Chars), false)
	}
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Right(row int, col int) *__Grid__ {
	grid.Grid[row-1][col-1].__UpdateDirection__("R")

	if col != grid.Size {
		grid.Grid[row-1][col] = grid.Grid[row-1][col-1]
		grid.Grid[row-1][col-1] = New("__Blank__", extras.Times(" ", grid.Chars), false)
	}
	grid.__UpdatePos__()

	return grid
}
func (grid *__Grid__) Print() *__Grid__ {
	for row := range grid.Grid {
		for col := range grid.Grid {
			fmt.Println(grid.Grid[row][col])
		}
	}

	return grid
}
func (grid *__Grid__) Return() *__Grid__ {
	fmt.Println("\n")

	return grid
}
func (grid *__Grid__) ToggleInvis() *__Grid__ {
	if grid.__Opacity__ == 0 {
		grid.__Opacity__ = 1
	} else {
		grid.__Opacity__ = 0
	}

	return grid
}

func Init(chars int, size int) *__Grid__ {
	var blankArray = make([][]*__Creature__, size)
	for i := range blankArray {
		blankArray[i] = make([]*__Creature__, size)
	}
	for row := range blankArray {
		for col := range blankArray {
			blankArray[row][col] = New("__Blank__", extras.Times(" ", chars), false)
		}
	}
	__CheckSize__(size)

	return &__Grid__{Grid: blankArray, Chars: chars, Size: size, __Opacity__: 1}
}
