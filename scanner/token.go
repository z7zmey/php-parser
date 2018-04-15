package scanner

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
)

// Token value returned by lexer
type Token struct {
	Value    string
	position *position.Position
	comments []*comment.Comment
}

// NewToken Token constructor
func NewToken(value []byte, pos *position.Position) *Token {
	return &Token{
		Value:    string(value),
		position: pos,
		comments: nil,
	}
}

func (t *Token) String() string {
	return string(t.Value)
}

// Position returns token position
func (t *Token) Position() *position.Position {
	return t.position
}

// Comments returns attached comments
func (t *Token) Comments() []*comment.Comment {
	return t.comments
}

// SetComments attach comments
func (t *Token) SetComments(comments []*comment.Comment) *Token {
	t.comments = comments
	return t
}
