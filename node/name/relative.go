package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Relative struct {
	Name
}

func NewRelative(Parts []node.Node) *Relative {
	return &Relative{
		Name{
			nil,
			nil,
			Parts,
		},
	}
}
