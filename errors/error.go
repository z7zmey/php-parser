package errors

import (
	"fmt"

	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/token"
)

// Error parsing error
type Error struct {
	Msg string
	Pos position.Position
}

// NewError creates and returns new Error
func NewError(msg string, t token.Token) *Error {
	return &Error{
		Msg: msg,
		Pos: position.Position{
			StartLine: t.StartLine,
			EndLine:   t.EndLine,
			StartPos:  t.StartPos,
			EndPos:    t.EndPos,
		},
	}
}

func (e *Error) String() string {
	return fmt.Sprintf("%s at line %d", e.Msg, e.Pos.StartLine)
}
