package binary_op

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type BinaryOp struct {
	node.SimpleNode
	left   node.Node
	right node.Node
}

func NewBinaryOp(left node.Node, right node.Node) node.Node {
	return BinaryOp{
		node.SimpleNode{Name: "BinaryOp", Attributes: make(map[string]string)},
		left,
		right,
	}
}

func (n BinaryOp) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.left != nil {
		fmt.Fprintf(out, "\n%vleft:", indent+"  ")
		n.left.Print(out, indent+"    ")
	}
	
	if n.right != nil {
		fmt.Fprintf(out, "\n%vright:", indent+"  ")
		n.right.Print(out, indent+"    ")
	}
}
