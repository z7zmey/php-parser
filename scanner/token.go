package scanner

import (
	"github.com/z7zmey/php-parser/comment"
)

// Token value returned by lexer
type Token struct {
	Value     string
	Comments  []*comment.Comment
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

func (t *Token) String() string {
	return string(t.Value)
}
