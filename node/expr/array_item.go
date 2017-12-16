package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ArrayItem struct {
	node.SimpleNode
	key node.Node
	val node.Node
	byRef bool
}

func NewArrayItem(key node.Node, val node.Node, byRef bool) node.Node {
	return ArrayItem{
		node.SimpleNode{Name: "ArrayItem", Attributes: make(map[string]string)},
		key,
		val,
		byRef,
	}
}

func (n ArrayItem) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vbyRef: %t", indent+"  ", n.byRef)

	if n.key != nil {
		fmt.Fprintf(out, "\n%vkey:", indent+"  ")
		n.key.Print(out, indent+"    ")
	}

	if n.val != nil {
		fmt.Fprintf(out, "\n%vval:", indent+"  ")
		n.val.Print(out, indent+"    ")
	}
}
