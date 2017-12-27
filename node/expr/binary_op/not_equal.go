package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n NotEqual) Name() string {
	return "NotEqual"
}

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(variable node.Node, expression node.Node) node.Node {
	return NotEqual{
		BinaryOp{
			"BinaryNotEqual",
			variable,
			expression,
		},
	}
}
