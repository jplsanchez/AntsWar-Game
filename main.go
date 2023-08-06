package main

import (
	"antswar/display"
)

func main() {
	displayer := display.ConsoleDisplay{}
	ui := display.NewUI(displayer)

	ui.Run()
}
