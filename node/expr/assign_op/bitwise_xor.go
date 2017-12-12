package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	AssignOp
}

func NewBitwiseXor(variable  node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		AssignOp{
			node.SimpleNode{Name: "AssignBitwiseXor", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
