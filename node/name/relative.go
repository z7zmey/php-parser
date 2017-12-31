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
			map[string]interface{}{},
			nil,
			parts,
		},
	}
}

func (n Relative) Name() string {
	return "Relative"
}

func (n Relative) Attributes() map[string]interface{} {
	return n.attributes
}
