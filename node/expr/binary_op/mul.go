package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	BinaryOp
}

func NewMul(variable  node.Node, expression node.Node) node.Node {
	return Mul{
		BinaryOp{
			node.SimpleNode{Name: "BinaryMul", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
