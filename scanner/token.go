package scanner

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
)

// Token value returned by lexer
type Token struct {
	Value        string
	FreeFloating []freefloating.String
	StartLine    int
	EndLine      int
	StartPos     int
	EndPos       int
}

func (t *Token) String() string {
	return string(t.Value)
}

func (t *Token) GetFreeFloatingToken() []freefloating.String {
	return []freefloating.String{
		{
			StringType: freefloating.TokenType,
			Value:      t.Value,
			Position: &position.Position{
				StartLine: t.StartLine,
				EndLine:   t.EndLine,
				StartPos:  t.StartPos,
				EndPos:    t.EndPos,
			},
		},
	}
}
