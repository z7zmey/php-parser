package scanner

import (
	"github.com/z7zmey/php-parser/pkg/token"
)

// Token value returned by lexer
type Token struct {
	Value     []byte
	Tokens    []token.Token
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}
