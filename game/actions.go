package game

import "errors"

type GameAction interface {
	DoAction()
	CanDo() error
}

type March struct{}

func (m March) DoAction()    {}
func (m March) CanDo() error { return errors.New("cannot march in this position") }

type Swap struct{}

func (s Swap) DoAction()    {}
func (s Swap) CanDo() error { return errors.New("cannot swap in this position") }

type Move struct{}

func (m Move) DoAction()    {}
func (m Move) CanDo() error { return errors.New("cannot move in this position") }

type Attack struct{}

func (a Attack) DoAction()    {}
func (a Attack) CanDo() error { return errors.New("cannot attack now or in this position") }
