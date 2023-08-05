package main

import (
	"fmt"
	"os"
)

const (
	ColorRed  = "\033[31m"
	ColorNone = "\033[0m"
)

type Displayer interface {
	Display(gb GameBoard)
}

type ConsoleDisplay struct{}

func (cd ConsoleDisplay) Display(gb GameBoard) {
	for i := 0; i < gb.Height(); i++ {
		cd.DisplayMiddle(i, gb)
		for j := 0; j < gb.Width(); j++ {
			cd.HandleAndDisplay(gb, i, j)
		}
		fmt.Println()
	}
}

func (ConsoleDisplay) DisplayMiddle(i int, gb GameBoard) {
	if i == gb.Height()/2 {
		fmt.Println("- - - - -")
	}
}

func (cd ConsoleDisplay) HandleAndDisplay(gb GameBoard, i int, j int) {
	val := gb[j][i].Value

	fmt.Fprintf(os.Stdout, "%s", ColorNone)
	if gb[j][i].Team == TeamRed {
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
