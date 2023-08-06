package game

import "errors"

type GameAction interface {
	DoAction()
	CanDo(pos Position, gb GameBoard, team Team) error
}

type March struct{}

func (m March) DoAction() { /*TODO*/ }
func (m March) CanDo(pos Position, gb GameBoard, team Team) error {
	hasFreeSpaces := false

	checkIfIsEmptySpaceAligned := func(i, j int) bool { return (i == pos.x || j == pos.y) && gb[i][j].Value == -1 }
	isEmptySpaceConnectedToPosition := func(i, j int) bool { return IsEmptySpaceConnectedToPosition(Position{i, j}, pos, gb, team) }

	gb.LoopThroughBoard(func(i, j int) {
		if checkIfIsEmptySpaceAligned(i, j) && isEmptySpaceConnectedToPosition(i, j) {
			hasFreeSpaces = true
		}
	})

	if !hasFreeSpaces {
		return errors.New("cannot march in this position because it has no free spaces in the row or column")
	}
	return nil
}

func IsEmptySpaceConnectedToPosition(emptySpace, pos Position, gb GameBoard, team Team) bool {

	sameColumn := emptySpace.x == pos.x
	sameRow := emptySpace.y == pos.y
	connectionIsNotValid := func(i, j int) bool {
		return gb[i][j].Value == -1 || (gb[i][j].Team == team.Enemy() && gb[i][j].Value >= 0)
	}

	if sameColumn && emptySpace.y > pos.y {
		for j := pos.y; j < emptySpace.y; j++ {
			if connectionIsNotValid(emptySpace.x, j) {
				return false
			}
		}
	}
	if sameColumn && emptySpace.y < pos.y {
		for j := pos.y; j > emptySpace.y; j-- {
			if connectionIsNotValid(emptySpace.x, j) {
				return false
			}
		}
	}
	if sameRow && emptySpace.x < pos.x {
		for i := pos.x; i > emptySpace.x; i-- {
			if connectionIsNotValid(i, emptySpace.y) {
				return false
			}
		}
	}
	if sameRow && emptySpace.x > pos.x {
		for i := pos.x; i < emptySpace.x; i++ {
			if connectionIsNotValid(i, emptySpace.y) {
				return false
			}
		}
	}
	return true
}

type Swap struct{}

func (s Swap) DoAction() { /*TODO*/ }
func (s Swap) CanDo(pos Position, gb GameBoard, team Team) error {
	validation := func(c Card) bool { return c.Value >= 0 && c.Team == team }
	hasAdjacentAllyCards := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentAllyCards {
		return errors.New("cannot swap in this position because it has no adjacent ally cards")
	}
	return nil
}

type Move struct{}

func (m Move) DoAction() { /*TODO*/ }
func (m Move) CanDo(pos Position, gb GameBoard, team Team) error {
	validation := func(c Card) bool { return c.Value < 0 }
	hasAdjacentFreeSpaces := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentFreeSpaces {
		return errors.New("cannot move in this position because it has no adjacent free spaces")
	}
	return nil
}

type Attack struct{}

func (a Attack) DoAction() { /*TODO*/ }
func (a Attack) CanDo(pos Position, gb GameBoard, team Team) error {
	validation := func(c Card) bool { return c.Value >= 0 && c.Team == team.Enemy() && c.Value < gb[pos.x][pos.y].Value }
	hasAdjacentEnemyCards := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentEnemyCards {
		return errors.New("cannot attack in this position because it has no adjacent enemy cards that can be attacked")
	}
	return nil
}

func ValidateAdjacentCards(gb GameBoard, pos Position, validate func(c Card) bool) bool {
	outOfRange := func(i, j int) bool {
		return i < 0 || i >= gb.Width() || j < 0 || j >= gb.Height()
	}

	if !outOfRange(pos.x+1, pos.y) {
		if c := gb[pos.x+1][pos.y]; validate(c) {
			return true
		}
	}
	if !outOfRange(pos.x-1, pos.y) {
		if c := gb[pos.x-1][pos.y]; validate(c) {
			return true
		}
	}
	if !outOfRange(pos.x, pos.y+1) {
		if c := gb[pos.x][pos.y+1]; validate(c) {
			return true
		}
	}
	if !outOfRange(pos.x, pos.y-1) {
		if c := gb[pos.x][pos.y-1]; validate(c) {
			return true
		}
	}
	return false
}
