package scanner

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
)

// Token value returned by lexer
type Token struct {
	Value    string
	Position *position.Position
	Comments []*comment.Comment
}

func (t *Token) String() string {
	return string(t.Value)
}
