package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Relative struct {
	Name
}

func NewRelative(parts []node.Node) node.Node {
	return Relative{
		Name{
			map[string]interface{}{},
			nil,
			parts,
		},
	}
}

func (n Relative) Attributes() map[string]interface{} {
	return n.attributes
}
