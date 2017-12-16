package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type ClassConstFetch struct {
	node.SimpleNode
	class node.Node
	name  token.Token
}

func NewClassConstFetch(class node.Node, name token.Token) node.Node {
	return ClassConstFetch{
		node.SimpleNode{Name: "ClassConstFetch", Attributes: make(map[string]string)},
		class,
		name,
	}
}

func (n ClassConstFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vname: %q", indent+"  ", n.name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}
}
