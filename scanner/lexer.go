// Package scanner transforms an input string into a stream of PHP tokens.
package scanner

import (
	"bufio"
	"bytes"
	"go/token"
	t "go/token"
	"io"
	"unicode"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"

	"github.com/cznic/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classUnicodeGraphic
	classOther
)

// Lval parsers yySymType must implement this interface
type Lval interface {
	Token(tkn *Token)
}

// Lexer php lexer
type Lexer struct {
	*lex.Lexer
	StateStack       []int
	PhpDocComment    string
	FreeFloating     []freefloating.String
	heredocLabel     string
	tokenBytesBuf    *bytes.Buffer
	TokenPool        *TokenPool
	WithFreeFloating bool
	lastToken        *Token
	Errors           []*errors.Error
}

// Rune2Class returns the rune integer id
func Rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}
	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}
	if unicode.IsGraphic(r) {
		return classUnicodeGraphic
	}
	if r == lex.RuneEOF {
		return int(r)
	}
	return classOther
}

func (l *Lexer) lexErrorFunc(p token.Pos, msg string) {
	pos := position.NewPosition(
		l.File.Line(p),
		l.File.Line(p),
		int(p),
		int(p),
	)
	l.Errors = append(l.Errors, errors.NewError(msg, pos))
}

// NewLexer the Lexer constructor
func NewLexer(src io.Reader, fName string) *Lexer {
	lexer := &Lexer{
		StateStack:    []int{0},
		tokenBytesBuf: &bytes.Buffer{},
		TokenPool:     &TokenPool{},
	}

	file := t.NewFileSet().AddFile(fName, -1, 1<<31-3)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(Rune2Class), lex.ErrorFunc(lexer.lexErrorFunc))
	if err != nil {
		panic(err)
	}
	lexer.Lexer = lx
	return lexer
}

func (l *Lexer) Error(msg string) {
	chars := l.Token()
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	pos := position.NewPosition(
		l.File.Line(firstChar.Pos()),
		l.File.Line(lastChar.Pos()),
		int(firstChar.Pos()),
		int(lastChar.Pos()),
	)

	l.Errors = append(l.Errors, errors.NewError(msg, pos))
}

func (l *Lexer) ungetChars(n int) []lex.Char {
	l.Unget(l.Lookahead())

	chars := l.Token()

	for i := 1; i <= n; i++ {
		char := chars[len(chars)-i]
		l.Unget(char)
	}

	buf := l.Token()
	buf = buf[:len(buf)-n]

	return buf
}

func (l *Lexer) pushState(state int) {
	l.StateStack = append(l.StateStack, state)
}

func (l *Lexer) popState() {
	len := len(l.StateStack)
	if len <= 1 {
		return
	}

	l.StateStack = l.StateStack[:len-1]
}

func (l *Lexer) Begin(state int) {
	len := len(l.StateStack)
	l.StateStack = l.StateStack[:len-1]
	l.StateStack = append(l.StateStack, state)
}

func (l *Lexer) getCurrentState() int {
	return l.StateStack[len(l.StateStack)-1]
}

func (l *Lexer) createToken(chars []lex.Char) *Token {
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	token := l.TokenPool.Get()
	token.FreeFloating = l.FreeFloating
	token.Value = l.tokenString(chars)

	// fmt.Println(l.tokenString(chars))

	token.StartLine = l.File.Line(firstChar.Pos())
	token.EndLine = l.File.Line(lastChar.Pos())
	token.StartPos = int(firstChar.Pos())
	token.EndPos = int(lastChar.Pos())

	return token
}

func (l *Lexer) tokenString(chars []lex.Char) string {
	l.tokenBytesBuf.Reset()

	for _, c := range chars {
		l.tokenBytesBuf.WriteRune(c.Rune)
	}

	return string(l.tokenBytesBuf.Bytes())
}

// free-floating

func (l *Lexer) addFreeFloating(t freefloating.StringType, chars []lex.Char) {
	if !l.WithFreeFloating {
		return
	}

	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	pos := position.NewPosition(
		l.File.Line(firstChar.Pos()),
		l.File.Line(lastChar.Pos()),
		int(firstChar.Pos()),
		int(lastChar.Pos()),
	)

	l.FreeFloating = append(l.FreeFloating, freefloating.String{
		StringType: t,
		Value:      l.tokenString(chars),
		Position:   pos,
	})
}
