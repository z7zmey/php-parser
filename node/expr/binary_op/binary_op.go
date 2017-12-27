package binary_op

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type BinaryOp struct {
	name  string
	left  node.Node
	right node.Node
}

func (n BinaryOp) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.left != nil {
		fmt.Fprintf(out, "\n%vleft:", indent+"  ")
		n.left.Print(out, indent+"    ")
	}

	if n.right != nil {
		fmt.Fprintf(out, "\n%vright:", indent+"  ")
		n.right.Print(out, indent+"    ")
	}
}
