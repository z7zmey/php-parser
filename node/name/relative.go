package name

import (
	"github.com/z7zmey/php-parser/node"
)

// Relative node
type Relative struct {
	Name
}

// NewRelative node constuctor
func NewRelative(Parts []node.Node) *Relative {
	return &Relative{
		Name{
			Parts,
		},
	}
}
