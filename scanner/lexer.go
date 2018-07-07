// Package scanner transforms an input string into a stream of PHP tokens.
package scanner

import (
	"bufio"
	"bytes"
	t "go/token"
	"io"
	"sync"
	"unicode"

	"github.com/z7zmey/php-parser/position"

	"github.com/cznic/golex/lex"
	"github.com/z7zmey/php-parser/meta"
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
	StateStack    []int
	PhpDocComment string
	Meta          []meta.Meta
	heredocLabel  string
	tokenBytesBuf *bytes.Buffer
	TokenPool     sync.Pool
	WithMeta      bool
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
	// return classOther
	return -1
}

// NewLexer the Lexer constructor
func NewLexer(src io.Reader, fName string) *Lexer {
	file := t.NewFileSet().AddFile(fName, -1, 1<<31-3)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(Rune2Class))
	if err != nil {
		panic(err)
	}

	return &Lexer{
		Lexer:         lx,
		StateStack:    []int{0},
		PhpDocComment: "",
		Meta:          nil,
		heredocLabel:  "",
		tokenBytesBuf: &bytes.Buffer{},
		TokenPool: sync.Pool{
			New: func() interface{} { return &Token{} },
		},
	}
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

func (l *Lexer) begin(state int) {
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

	token := l.TokenPool.Get().(*Token)
	token.Meta = l.Meta
	token.Value = l.tokenString(chars)

	token.StartLine = l.File.Line(firstChar.Pos())
	token.EndLine = l.File.Line(lastChar.Pos())
	token.StartPos = int(firstChar.Pos())
	token.EndPos = int(lastChar.Pos())

	return token
}

func (l *Lexer) addComments(chars []lex.Char) {
	if !l.WithMeta {
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

	c := meta.NewComment(l.tokenString(chars), pos)
	l.Meta = append(l.Meta, c)
}

func (l *Lexer) addWhiteSpace(chars []lex.Char) {
	if !l.WithMeta {
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

	c := meta.NewWhiteSpace(l.tokenString(chars), pos)
	l.Meta = append(l.Meta, c)
}

func (l *Lexer) tokenString(chars []lex.Char) string {
	l.tokenBytesBuf.Reset()

	for _, c := range chars {
		l.tokenBytesBuf.WriteRune(c.Rune)
	}

	return string(l.tokenBytesBuf.Bytes())
}
