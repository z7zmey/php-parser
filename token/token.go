package token

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
