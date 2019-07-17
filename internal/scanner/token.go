package scanner

// Token value returned by lexer
type Token struct {
	Type         TokenType
	HiddenTokens []Token
	StartLine    int
	EndLine      int
	StartPos     int
	EndPos       int
}
