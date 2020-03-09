package scanner

import (
	"bytes"
	"strings"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/version"
)

type Scanner interface {
	Lex(lval Lval) int
	ReturnTokenToPool(t *Token)
	GetPhpDocComment() string
	SetPhpDocComment(string)
	GetErrors() []*errors.Error
	GetWithFreeFloating() bool
	SetWithFreeFloating(bool)
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

	TokenPool        *TokenPool
	FreeFloating     []freefloating.String
	WithFreeFloating bool
	PhpDocComment    string
	lastToken        *Token
	Errors           []*errors.Error
	NewLines         NewLines
	PHPVersion       string
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

func (l *Lexer) GetWithFreeFloating() bool {
	return l.WithFreeFloating
}

func (l *Lexer) SetWithFreeFloating(b bool) {
	l.WithFreeFloating = b
}

func (l *Lexer) AddError(e *errors.Error) {
	l.Errors = append(l.Errors, e)
}

func (l *Lexer) SetErrors(e []*errors.Error) {
	l.Errors = e
}

func (lex *Lexer) setTokenPosition(token *Token) {
	token.StartLine = lex.NewLines.GetLine(lex.ts)
	token.EndLine = lex.NewLines.GetLine(lex.te - 1)
	token.StartPos = lex.ts
	token.EndPos = lex.te
}

func (lex *Lexer) addFreeFloating(t freefloating.StringType, ps, pe int) {
	if !lex.WithFreeFloating {
		return
	}

	pos := position.NewPosition(
		lex.NewLines.GetLine(lex.ts),
		lex.NewLines.GetLine(lex.te-1),
		lex.ts,
		lex.te,
	)

	lex.FreeFloating = append(lex.FreeFloating, freefloating.String{
		StringType: t,
		Value:      string(lex.data[ps:pe]),
		Position:   pos,
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
	r, err := version.Compare(lex.PHPVersion, "7.3")
	if err != nil {
		return lex.isHeredocEndSince73(p)
	}

	if r == -1 {
		return lex.isHeredocEndBefore73(p)
	}

	return lex.isHeredocEndSince73(p)
}

func (lex *Lexer) isHeredocEndBefore73(p int) bool {
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

func (lex *Lexer) isHeredocEndSince73(p int) bool {
	if lex.data[p-1] != '\r' && lex.data[p-1] != '\n' {
		return false
	}

	for lex.data[p] == ' ' || lex.data[p] == '\t' {
		p++
	}

	l := len(lex.heredocLabel)
	if len(lex.data) < p+l {
		return false
	}

	if len(lex.data) > p+l && isValidVarName(lex.data[p+l]) {
		return false
	}

	a := string(lex.heredocLabel)
	b := string(lex.data[p : p+l])

	_, _ = a, b

	if bytes.Equal(lex.heredocLabel, lex.data[p:p+l]) {
		lex.p = p
		return true
	}

	return false
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
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '_' || (r >= 0x80 && r <= 0xff)
}

func isValidVarName(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || (r >= 0x80 && r <= 0xff)
}
