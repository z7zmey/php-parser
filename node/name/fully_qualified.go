package name

import (
	"github.com/z7zmey/php-parser/node"
)

type FullyQualified struct {
	NameNode
}

func NewFullyQualified(parts []node.Node) node.Node {
	return FullyQualified{
		NameNode{
			"FullyQualifiedName",
			parts,
		},
	}
}

func (n FullyQualified) Name() string {
	return "FullyQualified"
}
