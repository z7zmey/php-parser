package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotIdentical struct {
	BinaryOp
}

func NewNotIdentical(variable node.Node, expression node.Node) node.Node {
	return NotIdentical{
		BinaryOp{
			node.SimpleNode{Name: "BinaryNotIdentical", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
