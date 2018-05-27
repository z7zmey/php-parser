package comment

import (
	"github.com/z7zmey/php-parser/position"
)

// Comment aggrigates information about comment /**
type Comment struct {
	value     string
	position  *position.Position
	tokenName TokenName
}

// NewComment - Comment constructor
func NewComment(value string, pos *position.Position) *Comment {
	return &Comment{
		value,
		pos,
		UnknownToken,
	}
}

// SetTokenName sets token name
func (c *Comment) SetTokenName(tokenName TokenName) {
	c.tokenName = tokenName
}

// TokenName returns token name
func (c *Comment) TokenName() TokenName {
	return c.tokenName
}

func (c *Comment) String() string {
	return c.value
}

// Position returns comment position
func (c *Comment) Position() *position.Position {
	return c.position
}
