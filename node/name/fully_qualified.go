package name

import (
	"github.com/z7zmey/php-parser/node"
)

type FullyQualified struct {
	Name
}

func NewFullyQualified(Parts []node.Node) *FullyQualified {
	return &FullyQualified{
		Name{
			nil,
			Parts,
		},
	}
}
