package main

import (
	"antswar/display"
	"antswar/display/console"
)

func main() {
	ui := display.NewUI(&console.ConsoleDisplay{})

	ui.Run()
}
