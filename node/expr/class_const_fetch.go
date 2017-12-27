package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ClassConstFetch) Name() string {
	return "ClassConstFetch"
}

type ClassConstFetch struct {
	name     string
	class    node.Node
	constant token.Token
}

func NewClassConstFetch(class node.Node, constant token.Token) node.Node {
	return ClassConstFetch{
		"ClassConstFetch",
		class,
		constant,
	}
}

func (n ClassConstFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)
	fmt.Fprintf(out, "\n%vname: %q", indent+"  ", n.constant.Value)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}
}
