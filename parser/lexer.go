package parser

import (
	"bufio"
	"bytes"
	"go/token"
	"io"
	"unicode"

	"github.com/cznic/golex/lex"
	"github.com/z7zmey/php-parser/comment"
	t "github.com/z7zmey/php-parser/token"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classUnicodeGraphic
	classOther
)

type lexer struct {
	*lex.Lexer
	stateStack    []int
	phpDocComment string
	comments      []comment.Comment
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
	if unicode.IsGraphic(r) {
		return classUnicodeGraphic
	}
	// return classOther
	return -1
}

func newLexer(src io.Reader, fName string) *lexer {
	file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}
	return &lexer{lx, []int{0}, "", nil}
}

func (l *lexer) ungetChars(n int) []lex.Char {
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

func (l *lexer) pushState(state int) {
	l.stateStack = append(l.stateStack, state)
}

func (l *lexer) popState() {
	len := len(l.stateStack)
	if len <= 1 {
		return
	}

	l.stateStack = l.stateStack[:len-1]
}

func (l *lexer) begin(state int) {
	len := len(l.stateStack)
	l.stateStack = l.stateStack[:len-1]
	l.stateStack = append(l.stateStack, state)
}

func (l *lexer) getCurrentState() int {
	return l.stateStack[len(l.stateStack)-1]
}

func (l *lexer) newToken(chars []lex.Char) t.Token {
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	startLine := l.File.Line(firstChar.Pos())
	endLine := l.File.Line(lastChar.Pos())
	startPos := int(firstChar.Pos())
	endPos := int(lastChar.Pos())

	return t.NewToken(l.charsToBytes(chars), startLine, endLine, startPos, endPos).SetComments(l.comments)
}

func (l *lexer) addComment(c comment.Comment) {
	l.comments = append(l.comments, c)
}

func (l *lexer) charsToBytes(chars []lex.Char) []byte {
	bytesBuf := bytes.Buffer{}

	for _, c := range chars {
		bytesBuf.WriteRune(c.Rune)
	}

	return bytesBuf.Bytes()
}
