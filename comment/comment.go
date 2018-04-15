package comment

import (
	"github.com/z7zmey/php-parser/position"
)

// Comment aggrigates information about comment /**
type Comment struct {
	value    string
	position *position.Position
}

// NewComment - Comment constructor
func NewComment(value string, pos *position.Position) *Comment {
	return &Comment{
		value,
		pos,
	}
}

func (c *Comment) String() string {
	return c.value
}

// Position returns comment position
func (c *Comment) Position() *position.Position {
	return c.position
}
