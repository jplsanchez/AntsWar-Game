package game

import (
	"errors"
)

type Position struct {
	x int
	y int
}

func (p Position) Validate(gb GameBoard) error {
	if p.x > gb.Width() || p.y > gb.Height() {
		return errors.New("Position out of bounds")
	}
	return nil
}

func (p Position) Sub(pos Position) Position {
	return Position{x: p.x - pos.x, y: p.y - pos.y}
}
func (p Position) Add(pos Position) Position {
	return Position{x: p.x + pos.x, y: p.y + pos.y}
}
func (p *Position) Normalize() {
	if p.x != 0 {
		p.x = p.x / Abs(p.x)
	}
	if p.y != 0 {
		p.y = p.y / Abs(p.y)
	}
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
