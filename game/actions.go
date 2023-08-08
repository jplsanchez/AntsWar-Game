package game

import "errors"

type GameAction interface {
	SetPosition(pos Position) GameAction
	SetTargetPosition(pos Position) GameAction
	SetGameBoard(gb *GameBoard) GameAction
	DoAction() error
	CanDo(pos Position, gb GameBoard, turn int) error
}

// MARCH

type March struct {
	pos       Position
	targetPos Position
	gb        *GameBoard
}

func (m *March) SetPosition(pos Position) GameAction       { m.pos = pos; return m }
func (m *March) SetTargetPosition(pos Position) GameAction { m.targetPos = pos; return m }
func (m *March) SetGameBoard(gb *GameBoard) GameAction     { m.gb = gb; return m }
func (m March) DoAction() error {
	if err := m.Validate(); err != nil {
		return err
	}

	vector := m.targetPos.Sub(m.pos)
	vector.Normalize()

	if vector.y == 0 && vector.x != 0 {
		for i := m.targetPos.x; i != m.pos.x; i -= vector.x {
			m.gb[i][m.pos.y] = m.gb[i-vector.x][m.pos.y]
		}
		m.gb[m.pos.x][m.pos.y] = EmptyCard()
	}

	if vector.x == 0 && vector.y != 0 {
		for j := m.targetPos.y; j != m.pos.y; j -= vector.y {
			m.gb[m.pos.x][j] = m.gb[m.pos.x][j-vector.y]
		}
		m.gb[m.pos.x][m.pos.y] = EmptyCard()
	}

	return nil
}

func (m March) CanDo(pos Position, gb GameBoard, turn int) error {
	hasFreeSpaces := false

	checkIfIsEmptySpaceAligned := func(i, j int) bool { return (i == pos.x || j == pos.y) && gb[i][j].Value == -1 }
	isEmptySpaceConnectedToPosition := func(i, j int) bool { return IsEmptySpaceConnectedToPosition(Position{i, j}, pos, gb) }

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
func (m March) Validate() error {
	if m.gb == nil {
		return errors.New("march is not valid because gameboard is nil")
	}
	if m.pos == m.targetPos {
		return errors.New("march is not valid because position and target position are the same")
	}
	if m.gb.GetCard(m.targetPos).Value != EmptyCard().Value {
		return errors.New("march is not valid because target position is not empty")
	}
	if err := m.pos.Validate(*m.gb); err != nil {
		return err
	}
	if vector := m.targetPos.Sub(m.pos); vector.x != 0 && vector.y != 0 {
		return errors.New("march is not valid because position and target position are not aligned")
	}
	if !IsEmptySpaceConnectedToPosition(m.targetPos, m.pos, *m.gb) {
		return errors.New("march is not valid because target position is not connected to position")
	}

	return nil
}

// SWAP

type Swap struct {
	pos       Position
	targetPos Position
	gb        *GameBoard
}

func (s *Swap) SetPosition(pos Position) GameAction       { s.pos = pos; return s }
func (s *Swap) SetTargetPosition(pos Position) GameAction { s.targetPos = pos; return s }
func (s *Swap) SetGameBoard(gb *GameBoard) GameAction     { s.gb = gb; return s }
func (s Swap) DoAction() error {
	if err := s.Validate(); err != nil {
		return err
	}
	card := s.gb.GetCard(s.pos)
	target := s.gb.GetCard(s.targetPos)

	s.gb.SetCard(s.pos, target)
	s.gb.SetCard(s.targetPos, card)
	return nil
}
func (s Swap) CanDo(pos Position, gb GameBoard, turn int) error {
	team := gb.GetCard(pos).Team
	validation := func(c Card) bool { return c.Value >= 0 && c.Team == team }
	hasAdjacentAllyCards := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentAllyCards {
		return errors.New("cannot swap in this position because it has no adjacent ally cards")
	}
	return nil
}
func (s Swap) Validate() error {
	if s.gb == nil {
		return errors.New("swap is not valid because gameboard is nil")
	}
	if s.pos == s.targetPos {
		return errors.New("swap is not valid because position and target position are the same")
	}
	if s.targetPos.Sub(s.pos).Magnitude() != 1 {
		return errors.New("swap is not valid because position and target position are not adjacent")
	}
	targetCard := s.gb.GetCard(s.targetPos)
	actionCard := s.gb.GetCard(s.pos)
	if targetCard.Value == EmptyCard().Value {
		return errors.New("swap is not valid because target position is empty")
	}
	if targetCard.Team != actionCard.Team {
		return errors.New("swap is not valid because target position is an ally card")
	}
	if err := s.pos.Validate(*s.gb); err != nil {
		return err
	}
	if err := s.targetPos.Validate(*s.gb); err != nil {
		return err
	}
	return nil
}

// MOVE

type Move struct {
	pos       Position
	targetPos Position
	gb        *GameBoard
}

func (m *Move) SetPosition(pos Position) GameAction       { m.pos = pos; return m }
func (m *Move) SetTargetPosition(pos Position) GameAction { m.targetPos = pos; return m }
func (m *Move) SetGameBoard(gb *GameBoard) GameAction     { m.gb = gb; return m }
func (m Move) DoAction() error {
	if err := m.Validate(); err != nil {
		return err
	}

	card := m.gb.GetCard(m.pos)
	m.gb.SetCard(m.targetPos, card)
	m.gb.SetCard(m.pos, EmptyCard())
	return nil
}
func (m Move) CanDo(pos Position, gb GameBoard, turn int) error {
	validation := func(c Card) bool { return c.Value < 0 }
	hasAdjacentFreeSpaces := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentFreeSpaces {
		return errors.New("cannot move in this position because it has no adjacent free spaces")
	}
	return nil
}
func (m Move) Validate() error {
	if m.gb == nil {
		return errors.New("move is not valid because gameboard is nil")
	}
	if m.pos == m.targetPos {
		return errors.New("move is not valid because position and target position are the same")
	}
	if err := m.pos.Validate(*m.gb); err != nil {
		return err
	}
	if err := m.targetPos.Validate(*m.gb); err != nil {
		return err
	}
	return nil
}

// ATTACK

type Attack struct {
	pos       Position
	targetPos Position
	gb        *GameBoard
}

func (a *Attack) SetPosition(pos Position) GameAction       { a.pos = pos; return a }
func (a *Attack) SetTargetPosition(pos Position) GameAction { a.targetPos = pos; return a }
func (a *Attack) SetGameBoard(gb *GameBoard) GameAction     { a.gb = gb; return a }

func (a Attack) CanDo(pos Position, gb GameBoard, turn int) error {
	if turn < 10 {
		return errors.New("cannot attack until played 5 times")
	}

	team := gb.GetCard(pos).Team
	validation := func(c Card) bool { return c.Value >= 0 && c.Team == team.Enemy() && c.Value <= gb[pos.x][pos.y].Value }
	hasAdjacentEnemyCards := ValidateAdjacentCards(gb, pos, validation)
	if !hasAdjacentEnemyCards {
		return errors.New("cannot attack in this position because it has no adjacent enemy cards that can be attacked")
	}
	return nil
}

func (a Attack) DoAction() error {
	if err := a.Validate(); err != nil {
		return err
	}

	card := a.gb.GetCard(a.pos)
	target := a.gb.GetCard(a.targetPos)

	if target.Value == None || target.Team != card.Team.Enemy() || target.Value > card.Value {
		return errors.New("cannot attack in this position because target cannot be attacked")
	}

	if target.Value == card.Value {
		a.gb.SetCard(a.targetPos, EmptyCard())
	} else {
		a.gb.SetCard(a.targetPos, card)
	}
	a.gb.SetCard(a.pos, EmptyCard())

	return nil
}

func (a Attack) Validate() error {
	if a.gb == nil {
		return errors.New("attack is not valid because gameboard is nil")
	}
	if a.pos == a.targetPos {
		return errors.New("attack is not valid because position and target position are the same")
	}
	if a.targetPos.Sub(a.pos).Magnitude() != 1 {
		return errors.New("attack is not valid because position and target position are not adjacent")
	}
	if err := a.pos.Validate(*a.gb); err != nil {
		return err
	}
	if err := a.targetPos.Validate(*a.gb); err != nil {
		return err
	}
	return nil
}

// HELPERS

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
func IsEmptySpaceConnectedToPosition(emptySpace, pos Position, gb GameBoard) bool {
	team := gb.GetCard(pos).Team
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
