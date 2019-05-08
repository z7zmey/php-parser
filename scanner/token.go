package scanner

type SkippedTokenType int8

const (
	SkippedTokenTypeWhitespace SkippedTokenType = iota
	SkippedTokenTypeComment
	SkippedTokenTypeToken
)

type SkippedToken struct {
	Type      SkippedTokenType
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

// Token value returned by lexer
type Token struct {
	SkippedTokens []SkippedToken
	StartLine     int
	EndLine       int
	StartPos      int
	EndPos        int
}

func (t *Token) AsSkippedTokens() []SkippedToken {
	return []SkippedToken{
		{
			Type:      SkippedTokenTypeToken,
			StartLine: t.StartLine,
			EndLine:   t.EndLine,
			StartPos:  t.StartPos,
			EndPos:    t.EndPos,
		},
	}
}
