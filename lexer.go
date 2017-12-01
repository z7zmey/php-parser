package main

import (
	"bufio"
	"go/token"
	"io"
	"unicode"

	"github.com/cznic/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

var sc int

type lexer struct {
	*lex.Lexer
}

var stateStack = []int{PHP}
var heredocLabel []byte

func pushState(state int) {
	sc = state
	stateStack = append(stateStack, state)
}

func popState() {
	len := len(stateStack)
	if len <= 1 {
		return
	}

	sc = stateStack[len-2]
	stateStack = stateStack[:len-1]
}

func begin(state int) {
	len := len(stateStack)
	stateStack = stateStack[:len-1]
	stateStack = append(stateStack, state)

	sc = state
}

func rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}
	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}
	// return classOther
	return -1
}

func newLexer(src io.Reader, dst io.Writer, fName string) *lexer {
	file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}
	return &lexer{lx}
}

func (l *lexer) unget(r rune) []byte {
	l.Unget(l.Lookahead())

	chars := l.Token()
	lastChar := chars[len(chars)-1]

	if lastChar.Rune != r {
		return l.TokenBytes(nil)
	}

	l.Unget(lastChar)

	buf := l.TokenBytes(nil)
	buf = buf[:len(buf)-1]

	return buf
}

func (l *lexer) ungetN(n int) []byte {
	l.Unget(l.Lookahead())

	chars := l.Token()

	for i := 1; i <= n; i++ {
		char := chars[len(chars)-i]
		l.Unget(char)
	}

	buf := l.TokenBytes(nil)
	buf = buf[:len(buf)-n]

	return buf
}
