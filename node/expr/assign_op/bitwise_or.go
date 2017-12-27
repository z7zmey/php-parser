package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseOr) Name() string {
	return "BitwiseOr"
}

type BitwiseOr struct {
	AssignOp
}

func NewBitwiseOr(variable node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		AssignOp{
			"AssignBitwiseOr",
			variable,
			expression,
		},
	}
}
