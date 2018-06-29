package meta

import (
	"github.com/z7zmey/php-parser/position"
)

// Comment aggrigates information about comment /**
type Comment struct {
	Value     string
	Position  *position.Position
	TokenName TokenName
}

// NewComment - Comment constructor
func NewComment(value string, pos *position.Position) *Comment {
	return &Comment{
		Value:     value,
		Position:  pos,
		TokenName: UnknownToken,
	}
}

// SetTokenName sets token name
func (c *Comment) SetTokenName(tokenName TokenName) {
	c.TokenName = tokenName
}

// GetTokenName returns token name
func (c *Comment) GetTokenName() TokenName {
	return c.TokenName
}

func (c *Comment) String() string {
	return c.Value
}

func (c *Comment) GetPosition() *position.Position {
	return c.Position
}
