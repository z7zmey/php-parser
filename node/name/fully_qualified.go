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
			map[string]interface{}{},
			nil,
			parts,
		},
	}
}

func (n FullyQualified) Attributes() map[string]interface{} {
	return n.attributes
}
