package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Coalesce) Name() string {
	return "Coalesce"
}

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(variable node.Node, expression node.Node) node.Node {
	return Coalesce{
		BinaryOp{
			"BinaryCoalesce",
			variable,
			expression,
		},
	}
}
