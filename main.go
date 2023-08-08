package main

import (
	"antswar/display"
)

func main() {
	ui := display.NewUI(&display.ConsoleDisplay{})

	ui.Run()
}
