package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(variable  node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		AssignOp{
			node.SimpleNode{Name: "AssignBitwiseAnd", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
