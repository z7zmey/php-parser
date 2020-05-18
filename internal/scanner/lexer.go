package scanner

import (
	"bytes"
	"strings"

	"github.com/z7zmey/php-parser/internal/version"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

type Scanner interface {
	Lex() *Token
	ReturnTokenToPool(t *Token)
}

type Config struct {
	WithHiddenTokens bool
	ErrHandlerFunc   func(*errors.Error)
}

type Lexer struct {
	data             []byte
	phpVersion       string
	withHiddenTokens bool
	errHandlerFunc   func(*errors.Error)

	p, pe, cs   int
	ts, te, act int
	stack       []int
	top         int

	heredocLabel []byte
	tokenPool    *TokenPool
	hiddenTokens []token.Token
	newLines     NewLines
}

func NewLexer(data []byte, phpVersion string, config Config) *Lexer {
	lex := &Lexer{
		data:             data,
		phpVersion:       phpVersion,
		errHandlerFunc:   config.ErrHandlerFunc,
		withHiddenTokens: config.WithHiddenTokens,

		pe:    len(data),
		stack: make([]int, 0),

		tokenPool: &TokenPool{},
		newLines:  NewLines{make([]int, 0, 128)},
	}

	initLexer(lex)

	return lex
}

func (lex *Lexer) ReturnTokenToPool(t *Token) {
	lex.tokenPool.Put(t)
}

func (lex *Lexer) setTokenPosition(token *Token) {
	token.Position.StartLine = lex.newLines.GetLine(lex.ts)
	token.Position.EndLine = lex.newLines.GetLine(lex.te - 1)
	token.Position.StartPos = lex.ts
	token.Position.EndPos = lex.te
}

func (lex *Lexer) addHiddenToken(id TokenID, ps, pe int) {
	if !lex.withHiddenTokens {
		return
	}

	lex.hiddenTokens = append(lex.hiddenTokens, token.Token{
		ID:    token.ID(id),
		Value: lex.data[ps:pe],
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
	r, err := version.Compare(lex.phpVersion, "7.3")
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

func (lex *Lexer) error(msg string) {
	if lex.errHandlerFunc == nil {
		return
	}

	pos := position.NewPosition(
		lex.newLines.GetLine(lex.ts),
		lex.newLines.GetLine(lex.te-1),
		lex.ts,
		lex.te,
	)

	lex.errHandlerFunc(errors.NewError(msg, pos))
}

func isValidVarNameStart(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '_' || (r >= 0x80 && r <= 0xff)
}

func isValidVarName(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || (r >= 0x80 && r <= 0xff)
}
