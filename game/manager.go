package game

import "errors"

type Position struct {
	x int
	y int
}

type GameStage uint8

const (
	StartTurnStage GameStage = iota
	ChooseCardStage
	ChooseActionStage
	ResolveActionStage
	EndTurnStage
)

type States struct {
	Stage          GameStage
	Team           Team
	CardSelected   Position
	GameFinished   bool
	ActionSelected GameAction
}

type GameManager struct {
	Turns  int
	Board  *GameBoard
	States *States
}

func NewGameManager(startingTeam Team, board *GameBoard) *GameManager {
	states := States{Team: startingTeam}
	gm := GameManager{Turns: 0, Board: board, States: &states}
	return &gm
}

func (gm *GameManager) StartPlayersTurn() {
	switch gm.States.Team {
	case TeamBlack:
		gm.States.Team = TeamRed
	case TeamRed:
		gm.States.Team = TeamBlack
	}
}

func (gm *GameManager) ChooseCard(x, y int) error {
	if x >= gm.Board.Width() || y >= gm.Board.Width() {
		return errors.New("x or y out of dimensions of the board")
	}
	gm.States.CardSelected = Position{x: x, y: y}
	return nil
}

func (gm *GameManager) ChooseAction(action GameAction) error {
	if err := action.CanDo(); err != nil {
		return err
	}
	gm.States.ActionSelected = action
	return nil
}

func (gm GameManager) ResolveAction() {
	gm.States.ActionSelected.DoAction()
}

func (gm *GameManager) CheckIfGameFinished() {
	if gm.Board.CountQueens() < 2 {
		gm.States.GameFinished = true
	}
}
