package main

import (
	"antswar/display"
	"antswar/display/sockets"
)

func main() {
	// DI := &console.ConsoleDisplay{}
	DI := &sockets.SocketDisplay{}
	ui := display.NewUI(DI)

	ui.Run()
}
