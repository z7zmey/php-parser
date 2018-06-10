package errors

import (
	"fmt"

	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

// Error parsing error
type Error struct {
	Msg string
	Pos *position.Position
}

// NewError creates and returns new Error
func NewError(msg string, t *scanner.Token) *Error {
	return &Error{
		Msg: msg,
		Pos: t.Position,
	}
}

func (e *Error) String() string {
	return fmt.Sprintf("%s at line %d", e.Msg, e.Pos.StartLine)
}
