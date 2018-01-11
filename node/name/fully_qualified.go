package name

import (
	"github.com/z7zmey/php-parser/node"
)

// FullyQualified node
type FullyQualified struct {
	Name
}

// NewFullyQualified node constuctor
func NewFullyQualified(Parts []node.Node) *FullyQualified {
	return &FullyQualified{
		Name{
			Parts,
		},
	}
}
