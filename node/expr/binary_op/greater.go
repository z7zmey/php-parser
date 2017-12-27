package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Greater) Name() string {
	return "Greater"
}

type Greater struct {
	BinaryOp
}

func NewGreater(variable node.Node, expression node.Node) node.Node {
	return Greater{
		BinaryOp{
			"BinaryGreater",
			variable,
			expression,
		},
	}
}
