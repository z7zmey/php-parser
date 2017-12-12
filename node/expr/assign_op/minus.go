package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	AssignOp
}

func NewMinus(variable  node.Node, expression node.Node) node.Node {
	return Minus{
		AssignOp{
			node.SimpleNode{Name: "AssignMinus", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
