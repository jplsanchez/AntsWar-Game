package sockets

import (
	"antswar/game"
	"encoding/json"
)

type Type string

const (
	GameBoardType  Type = "GameBoard"
	StringType     Type = "string"
	StringListType Type = "...string"
	TeamType       Type = "Team"
	PositionType   Type = "[]int{x, y}"
)

type SocketDisplay struct {
	server *Server
}

type SocketDTO struct {
	Name Type `json:"name"`
	Data any  `json:"data"`
}

func (sd *SocketDisplay) Init() {
	sd.server = NewServer(":8080")
}
func (sd SocketDisplay) UpdateBoard(gb game.GameBoard) {
	json, err := json.Marshal(SocketDTO{GameBoardType, gb})
	if err != nil {
		panic(err)
	}
	_, err = sd.server.Broadcast(json)
	if err != nil {
		panic(err)
	}
}
func (sd SocketDisplay) DisplayMessage(message string) {
	json, err := json.Marshal(SocketDTO{StringType, message})
	if err != nil {
		panic(err)
	}
	_, err = sd.server.Broadcast(json)
	if err != nil {
		panic(err)
	}
}
func (sd SocketDisplay) AskForString(messages ...string) (string, error) {
	json, err := json.Marshal(SocketDTO{StringListType, messages})
	if err != nil {
		return "", err
	}
	response, err := sd.server.Broadcast(json)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
func (sd SocketDisplay) SetPlayer(player game.Team) {
	json, err := json.Marshal(SocketDTO{TeamType, player})
	if err != nil {
		panic(err)
	}
	_, err = sd.server.Broadcast(json)
	if err != nil {
		panic(err)
	}
}
func (sd SocketDisplay) SetHighlight(x, y int) {
	json, err := json.Marshal(SocketDTO{PositionType, []int{x, y}})
	if err != nil {
		panic(err)
	}
	_, err = sd.server.Broadcast(json)
	if err != nil {
		panic(err)
	}
}
