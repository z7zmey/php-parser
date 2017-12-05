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
			node.SimpleNode{Name: "RelativeName", Attributes: make(map[string]string)},
			parts,
		},
	}
}

