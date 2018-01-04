package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Relative struct {
	Name
}

func NewRelative(Parts []node.Node) node.Node {
	return &Relative{
		Name{
			nil,
			Parts,
		},
	}
}
