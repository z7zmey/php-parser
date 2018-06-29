package meta

import (
	"github.com/z7zmey/php-parser/position"
)

type WhiteSpace struct {
	Value     string
	Position  *position.Position
	TokenName TokenName
}

func NewWhiteSpace(value string, pos *position.Position) *WhiteSpace {
	return &WhiteSpace{
		Value:     value,
		Position:  pos,
		TokenName: UnknownToken,
	}
}

// SetTokenName sets token name
func (c *WhiteSpace) SetTokenName(tokenName TokenName) {
	c.TokenName = tokenName
}

// GetTokenName returns token name
func (c *WhiteSpace) GetTokenName() TokenName {
	return c.TokenName
}

func (el *WhiteSpace) String() string {
	return el.Value
}

func (el *WhiteSpace) GetPosition() *position.Position {
	return el.Position
}
