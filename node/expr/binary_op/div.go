package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	BinaryOp
}

func NewDiv(variable  node.Node, expression node.Node) node.Node {
	return Div{
		BinaryOp{
			node.SimpleNode{Name: "BinaryDiv", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
