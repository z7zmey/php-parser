package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseAnd) Name() string {
	return "BitwiseAnd"
}

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		AssignOp{
			"AssignBitwiseAnd",
			variable,
			expression,
		},
	}
}
