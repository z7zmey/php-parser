package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	AssignOp
}

func NewPow(variable  node.Node, expression node.Node) node.Node {
	return Pow{
		AssignOp{
			node.SimpleNode{Name: "AssignPow", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
