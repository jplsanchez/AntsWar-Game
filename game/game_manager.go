package game

import (
	"errors"
)

type GameStage uint8

const (
	StartTurnStage GameStage = iota
	ChooseCardStage
	ChooseActionStage
	ResolveActionStage
	EndTurnStage
)

type GameManager struct {
	Turns  int
	Board  *GameBoard
	States *States
}

type States struct {
	Stage            GameStage
	Team             Team
	SelectedPosition Position
	GameFinished     bool
	ActionSelected   GameAction
}

func NewGameManager(startingTeam Team, board *GameBoard) *GameManager {
	states := States{
		Team:         startingTeam,
		Stage:        ChooseCardStage,
		GameFinished: false,
	}
	gm := GameManager{Turns: 10, Board: board, States: &states}
	return &gm
}

func (gm *GameManager) NextGameStage() GameStage {
	if gm.States.Stage == EndTurnStage {
		gm.States.Stage = StartTurnStage
		return gm.States.Stage
	}
	gm.States.Stage++
	return gm.States.Stage
}

func (gm *GameManager) StartTurn() {
	gm.Turns++
	gm.States.Team = gm.States.Team.Enemy()
}

func (gm *GameManager) ChooseCard(x, y int) error {
	if y >= gm.Board.Height() || x >= gm.Board.Width() {
		return errors.New("x or y out of dimensions of the board")
	}
	card := gm.Board.GetCard(Position{x: x, y: y})
	if card.Value == None {
		return errors.New("no card in that position")
	}
	if card.Team != gm.States.Team {
		return errors.New("card is not from the team of the player")
	}

	gm.States.SelectedPosition = Position{x: x, y: y}
	return nil
}

func (gm *GameManager) ChooseAction(action GameAction, pos Position) error {
	if err := action.CanDo(pos, *gm.Board, gm.Turns); err != nil {
		return err
	}
	gm.States.ActionSelected = action
	return nil
}

func (gm GameManager) ResolveAction(x, y int) error {
	gm.States.ActionSelected.SetPosition(gm.States.SelectedPosition)
	gm.States.ActionSelected.SetTargetPosition(Position{x: x, y: y})
	gm.States.ActionSelected.SetGameBoard(gm.Board)

	return gm.States.ActionSelected.DoAction()
}

func (gm *GameManager) CheckIfGameFinished() string {
	queens := gm.Board.CountQueens()
	if len(queens) == 0 {
		gm.States.GameFinished = true
		return "Draw"
	}
	if len(queens) == 1 {
		gm.States.GameFinished = true
		return queens[0].String() + " won!"
	}
	return ""
}
