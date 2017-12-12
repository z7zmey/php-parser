package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	AssignOp
}

func NewMul(variable  node.Node, expression node.Node) node.Node {
	return Mul{
		AssignOp{
			node.SimpleNode{Name: "AssignMul", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
