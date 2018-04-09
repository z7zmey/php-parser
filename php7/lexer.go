package php7

import (
	"bufio"
	goToken "go/token"
	"io"

	"github.com/cznic/golex/lex"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/scanner"
	"github.com/z7zmey/php-parser/token"
)

type lexer struct {
	scanner.Lexer
	lastToken *token.Token
	errors    []*errors.Error
}

func (l *lexer) Lex(lval *yySymType) int {
	t := l.Lexer.Lex(lval)
	l.lastToken = &lval.token

	return t
}

func (l *lexer) Error(msg string) {
	l.errors = append(l.errors, errors.NewError(msg, *l.lastToken))
}

func (lval *yySymType) Token(t token.Token) {
	lval.token = t
}

func newLexer(src io.Reader, fName string) *lexer {
	file := goToken.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(scanner.Rune2Class))
	if err != nil {
		panic(err)
	}

	scanner := scanner.Lexer{
		Lexer:         lx,
		StateStack:    []int{0},
		PhpDocComment: "",
		Comments:      nil,
	}

	return &lexer{
		scanner,
		nil,
		nil,
	}
}
