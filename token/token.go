package token

import (
	"github.com/z7zmey/php-parser/comment"
)

type TokenInterface interface {
	GetValue() string
	GetStartLine() int
	GetEndLine() int
	Comments() []comment.Comment
	SetComments(comments []comment.Comment) Token
}

type Token struct {
	Value     string
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
	comments  []comment.Comment
}

func NewToken(value []byte, startLine int, endLine int, startPos int, endPos int) Token {
	return Token{string(value), startLine, endLine, startPos, endPos, nil}
}

func (t Token) String() string {
	return string(t.Value)
}

func (t Token) GetValue() string {
	return t.Value
}
func (t Token) GetStartLine() int {
	return t.StartLine
}
func (t Token) GetEndLine() int {
	return t.EndLine
}

func (t Token) Comments() []comment.Comment {
	return t.comments
}

func (t Token) SetComments(comments []comment.Comment) Token {
	t.comments = comments
	return t
}
