package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Relative struct {
	NameNode
}

func NewRelative(parts []node.Node) node.Node {
	return Relative{
		NameNode{
			"RelativeName",
			parts,
		},
	}
}

func (n Relative) Name() string {
	return "Relative"
}
