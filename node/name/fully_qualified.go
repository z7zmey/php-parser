package name

import (
	"github.com/z7zmey/php-parser/node"
)

type FullyQualified struct {
	Name
}

func NewFullyQualified(parts []node.Node) node.Node {
	return FullyQualified{
		Name{
			node.SimpleNode{Name: "FullyQualifiedName", Attributes: make(map[string]string)},
			parts,
		},
	}
}
