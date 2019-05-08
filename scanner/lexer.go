package scanner

import (
	"bytes"
	"strings"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/position"
)

type Scanner interface {
	Reset(data []byte)
	Lex(lval Lval) int
	ReturnTokenToPool(t *Token)
	GetPhpDocComment() string
	SetPhpDocComment(string)
	GetErrors() []*errors.Error
	AddError(e *errors.Error)
	SetErrors(e []*errors.Error)
}

// Lval parsers yySymType must implement this interface
type Lval interface {
	Token(tkn *Token)
}

type Lexer struct {
	data         []byte
	p, pe, cs    int
	ts, te, act  int
	stack        []int
	top          int
	heredocLabel []byte

	TokenPool     *TokenPool
	PhpDocComment string
	lastToken     *Token
	Errors        []*errors.Error
	NewLines      NewLines
}

func (l *Lexer) Reset(data []byte) {
	l.data = data
	l.pe = len(data)
	l.p = 0

	l.cs = lexer_start
	l.top = 0
	l.ts = 0
	l.te = 0
	l.act = 0
	l.stack = l.stack[:0]

	l.NewLines.Reset()
}

func (l *Lexer) ReturnTokenToPool(t *Token) {
	l.TokenPool.Put(t)
}

func (l *Lexer) GetPhpDocComment() string {
	return l.PhpDocComment
}

func (l *Lexer) SetPhpDocComment(s string) {
	l.PhpDocComment = s
}

func (l *Lexer) GetErrors() []*errors.Error {
	return l.Errors
}

func (l *Lexer) AddError(e *errors.Error) {
	l.Errors = append(l.Errors, e)
}

func (l *Lexer) SetErrors(e []*errors.Error) {
	l.Errors = e
}

func (lex *Lexer) prepareToken(token *Token) *Token {
	token.StartLine = lex.NewLines.GetLine(lex.ts)
	token.EndLine = lex.NewLines.GetLine(lex.te - 1)
	token.StartPos = lex.ts
	token.EndPos = lex.te

	return token
}

func (lex *Lexer) addSkippedToken(token *Token, t SkippedTokenType, ps, pe int) {
	token.SkippedTokens = append(token.SkippedTokens, SkippedToken{
		Type:      t,
		StartLine: lex.NewLines.GetLine(lex.ts),
		EndLine:   lex.NewLines.GetLine(lex.te - 1),
		StartPos:  lex.ts,
		EndPos:    lex.te,
	})
}

func (lex *Lexer) isNotStringVar() bool {
	p := lex.p
	if lex.data[p-1] == '\\' && lex.data[p-2] != '\\' {
		return true
	}

	if len(lex.data) < p+1 {
		return true
	}

	if lex.data[p] == '$' && (lex.data[p+1] == '{' || isValidVarNameStart(lex.data[p+1])) {
		return false
	}

	if lex.data[p] == '{' && lex.data[p+1] == '$' {
		return false
	}

	return true
}

func (lex *Lexer) isNotStringEnd(s byte) bool {
	p := lex.p
	if lex.data[p-1] == '\\' && lex.data[p-2] != '\\' {
		return true
	}

	return !(lex.data[p] == s)
}

func (lex *Lexer) isHeredocEnd(p int) bool {
	if lex.data[p-1] != '\r' && lex.data[p-1] != '\n' {
		return false
	}

	l := len(lex.heredocLabel)
	if len(lex.data) < p+l {
		return false
	}

	if len(lex.data) > p+l && lex.data[p+l] != ';' && lex.data[p+l] != '\r' && lex.data[p+l] != '\n' {
		return false
	}

	if len(lex.data) > p+l+1 && lex.data[p+l] == ';' && lex.data[p+l+1] != '\r' && lex.data[p+l+1] != '\n' {
		return false
	}

	return bytes.Equal(lex.heredocLabel, lex.data[p:p+l])
}

func (lex *Lexer) isNotHeredocEnd(p int) bool {
	return !lex.isHeredocEnd(p)
}

func (lex *Lexer) growCallStack() {
	if lex.top == len(lex.stack) {
		lex.stack = append(lex.stack, 0)
	}
}

func (lex *Lexer) isNotPhpCloseToken() bool {
	if lex.p+1 == len(lex.data) {
		return true
	}

	return lex.data[lex.p] != '?' || lex.data[lex.p+1] != '>'
}

func (lex *Lexer) isNotNewLine() bool {
	if lex.data[lex.p] == '\n' && lex.data[lex.p-1] == '\r' {
		return true
	}

	return lex.data[lex.p-1] != '\n' && lex.data[lex.p-1] != '\r'
}

func (lex *Lexer) call(state int, fnext int) {
	lex.growCallStack()

	lex.stack[lex.top] = state
	lex.top++

	lex.p++
	lex.cs = fnext
}

func (lex *Lexer) ret(n int) {
	lex.top = lex.top - n
	if lex.top < 0 {
		lex.top = 0
	}
	lex.cs = lex.stack[lex.top]
	lex.p++
}

func (lex *Lexer) ungetStr(s string) {
	tokenStr := string(lex.data[lex.ts:lex.te])
	if strings.HasSuffix(tokenStr, s) {
		lex.ungetCnt(len(s))
	}
}

func (lex *Lexer) ungetCnt(n int) {
	lex.p = lex.p - n
	lex.te = lex.te - n
}

func (lex *Lexer) Error(msg string) {
	pos := position.NewPosition(
		lex.NewLines.GetLine(lex.ts),
		lex.NewLines.GetLine(lex.te-1),
		lex.ts,
		lex.te,
	)

	lex.Errors = append(lex.Errors, errors.NewError(msg, pos))
}

func isValidVarNameStart(r byte) bool {
	return r >= 'A' && r <= 'Z' || r == '_' || r >= 'a' && r <= 'z' || r >= '\u007f' && r <= 'ÿ'
}
