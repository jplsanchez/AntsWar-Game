package sockets

import (
	"antswar/game"
	"encoding/json"
	"log"
)

type Type string

const (
	GameBoardType  Type = "GameBoard"
	StringType     Type = "SingleString"
	StringListType Type = "StringCollection"
	TeamType       Type = "Team"
	PositionType   Type = "PositionArray"
)

type SocketDisplay struct {
	server *Server
}

type SocketDTO struct {
	Name Type `json:"name"`
	Data any  `json:"data"`
}

func (sd *SocketDisplay) Init() {
	ready := make(chan bool, 1)
	log.Println("1")
	go sd.StartServer(ready)
	log.Println("2")

	<-ready

	log.Println("3")
}

func (sd *SocketDisplay) StartServer(ready chan bool) {
	sd.server = NewServer(":8080", ready)
	sd.server.Start()
}
func (sd SocketDisplay) UpdateBoard(gb game.GameBoard) {
	json, err := json.Marshal(SocketDTO{GameBoardType, gb})
	if err != nil {
		log.Fatal(err)
	}
	_, err = sd.server.SendRequest(json)
	if err != nil {
		log.Fatal(err)
	}
}
func (sd SocketDisplay) DisplayMessage(message string) {
	json, err := json.Marshal(SocketDTO{StringType, message})
	if err != nil {
		log.Fatal(err)
	}
	_, err = sd.server.SendRequest(json)
	if err != nil {
		log.Fatal(err)
	}
}
func (sd SocketDisplay) AskForString(messages ...string) (string, error) {
	json, err := json.Marshal(SocketDTO{StringListType, messages})
	if err != nil {
		return "", err
	}
	response, err := sd.server.SendRequest(json)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
func (sd SocketDisplay) SetPlayer(player game.Team) {
	json, err := json.Marshal(SocketDTO{TeamType, player})
	if err != nil {
		log.Fatal(err)
	}
	_, err = sd.server.SendRequest(json)
	if err != nil {
		log.Fatal(err)
	}
}
func (sd SocketDisplay) SetHighlight(x, y int) {
	json, err := json.Marshal(SocketDTO{PositionType, []int{x, y}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = sd.server.SendRequest(json)
	if err != nil {
		log.Fatal(err)
	}
}
