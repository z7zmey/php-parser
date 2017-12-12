package assign_op

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Assign struct {
	AssignOp
	byRef bool
}

func NewAssign(variable  node.Node, expression node.Node, byRef bool) node.Node {
	return Assign{
		AssignOp{
			node.SimpleNode{Name: "AssignAssign", Attributes: make(map[string]string)},
			variable,
			expression,
		},
		byRef,
	}
}

func (n Assign) Print(out io.Writer, indent string) {
	n.AssignOp.Print(out, indent)
	
	fmt.Fprintf(out, "\n%vbyRef: %t", indent+"  ", n.byRef)
}
