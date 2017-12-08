package token

type TokenInterface interface {
	GetValue() string
	GetStartLine() int
	GetEndLine() int
}

type Token struct {
	Value     string
	StartLine int
	EndLine   int
}

func NewToken(value []byte, startLine int, endLine int) Token {
	return Token{string(value), startLine, endLine}
}

func (t Token) String() string {
	return string(t.Value)
}

func (t Token) GetValue() string {
	return t.Value
}
func (t Token) GetStartLine() int {
	return t.StartLine
}
func (t Token) GetEndLine() int {
	return t.EndLine
}
