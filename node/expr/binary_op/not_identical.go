package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n NotIdentical) Name() string {
	return "NotIdentical"
}

type NotIdentical struct {
	BinaryOp
}

func NewNotIdentical(variable node.Node, expression node.Node) node.Node {
	return NotIdentical{
		BinaryOp{
			"BinaryNotIdentical",
			variable,
			expression,
		},
	}
}
