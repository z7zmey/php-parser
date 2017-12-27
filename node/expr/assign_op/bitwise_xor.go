package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	AssignOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		AssignOp{
			"AssignBitwiseXor",
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Name() string {
	return "BitwiseXor"
}
