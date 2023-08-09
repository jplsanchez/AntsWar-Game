package console

import "github.com/fatih/color"

const MessageLogSize = 16

type MessageLog []Message

type Message struct {
	message   string
	textColor color.Color
}

func (ml MessageLog) Print() {
	for i := 0; i < len(ml); i++ {
		color := ml[i].textColor
		color.Println(ml[i].message)
	}
}

func (ml *MessageLog) Append(message string, textColor color.Color) *MessageLog {
	slice := *ml
	slice = append(slice, Message{message, textColor})
	if len(slice) > MessageLogSize {
		slice = slice[1:]
	}
	ml = &slice
	return ml
}
