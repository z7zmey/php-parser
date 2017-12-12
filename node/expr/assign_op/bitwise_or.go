package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	AssignOp
}

func NewBitwiseOr(variable  node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		AssignOp{
			node.SimpleNode{Name: "AssignBitwiseOr", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
