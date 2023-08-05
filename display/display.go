package display

import (
	"antswar/game"
	"fmt"
	"os"
)

const (
	ColorRed  = "\033[31m"
	ColorNone = "\033[0m"
)

type Displayer interface {
	Display(gb game.GameBoard)
}

type ConsoleDisplay struct{}

func (cd ConsoleDisplay) Display(gb game.GameBoard) {
	for i := 0; i < gb.Height(); i++ {
		cd.DisplayMiddle(i, gb)
		for j := 0; j < gb.Width(); j++ {
			cd.HandleAndDisplay(gb, i, j)
		}
		fmt.Println()
	}
}

func (ConsoleDisplay) DisplayMiddle(i int, gb game.GameBoard) {
	if i == gb.Height()/2 {
		fmt.Println("- - - - -")
	}
}

func (cd ConsoleDisplay) HandleAndDisplay(gb game.GameBoard, i int, j int) {
	val := gb[j][i].Value

	fmt.Fprintf(os.Stdout, "%s", ColorNone)
	if gb[j][i].Team == game.TeamRed {
		fmt.Fprintf(os.Stdout, "%s", ColorRed)
	}

	if val >= 1 && val <= 9 {
		fmt.Printf("%d ", gb[j][i].Value)
		return
	}

	if val == 0 {
		fmt.Printf("Q ")
		return
	}

	fmt.Printf("  ")
	fmt.Fprintf(os.Stdout, "%s", ColorNone)
}
