package game

import "errors"

type GameStage uint8

const (
	StartTurnStage GameStage = iota
	ChooseCardStage
	ChooseActionStage
	ResolveActionStage
	EndTurnStage
)

type States struct {
	Stage            GameStage
	Team             Team
	SelectedPosition Position
	GameFinished     bool
	ActionSelected   GameAction
}

type GameManager struct {
	Turns  int
	Board  *GameBoard
	States *States
}

func NewGameManager(startingTeam Team, board *GameBoard) *GameManager {
	states := States{
		Team:         startingTeam,
		Stage:        ChooseCardStage,
		GameFinished: false,
	}
	gm := GameManager{Turns: 0, Board: board, States: &states}
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

func (gm *GameManager) StartPlayersTurn() {
	gm.States.Team.Enemy()
}

func (gm *GameManager) ChooseCard(x, y int) error {
	if x >= gm.Board.Width() || y >= gm.Board.Width() {
		return errors.New("x or y out of dimensions of the board")
	}
	gm.States.SelectedPosition = Position{x: x, y: y}
	return nil
}

func (gm *GameManager) ChooseAction(action GameAction, pos Position) error {
	if err := action.CanDo(pos, *gm.Board); err != nil {
		return err
	}
	gm.States.ActionSelected = action
	return nil
}

func (gm GameManager) ResolveAction(target Position) {
	gm.States.ActionSelected.SetPosition(gm.States.SelectedPosition)
	gm.States.ActionSelected.SetTargetPosition(target)
	gm.States.ActionSelected.SetGameBoard(gm.Board)

	gm.States.ActionSelected.DoAction()
}

func (gm *GameManager) CheckIfGameFinished() {
	if gm.Board.CountQueens() < 2 {
		gm.States.GameFinished = true
	}
}
