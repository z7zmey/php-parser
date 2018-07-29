package meta

import (
	"github.com/z7zmey/php-parser/position"
)

// Data contain additional information that isn't part of AST
type Data struct {
	Value     string
	Type      Type
	Position  *position.Position
	TokenName TokenName
}

func (d *Data) String() string {
	return d.Value
}
