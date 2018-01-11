package token

import (
	"github.com/z7zmey/php-parser/comment"
)

// Token value returned by lexer
type Token struct {
	Value     string
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
	comments  []comment.Comment
}

// NewToken Token constructor
// TODO: return pointer
func NewToken(value []byte, startLine int, endLine int, startPos int, endPos int) Token {
	return Token{string(value), startLine, endLine, startPos, endPos, nil}
}

func (t Token) String() string {
	return string(t.Value)
}

// Comments returns attached comments
func (t Token) Comments() []comment.Comment {
	return t.comments
}

// SetComments attach comments
func (t Token) SetComments(comments []comment.Comment) Token {
	t.comments = comments
	return t
}
