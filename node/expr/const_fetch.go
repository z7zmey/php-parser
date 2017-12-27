package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n ConstFetch) Name() string {
	return "ConstFetch"
}

type ConstFetch struct {
	name     string
	constant node.Node
}

func NewConstFetch(constant node.Node) node.Node {
	return ConstFetch{
		"ConstFetch",
		constant,
	}
}

func (n ConstFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.constant != nil {
		fmt.Fprintf(out, "\n%vconstant:", indent+"  ")
		n.constant.Print(out, indent+"    ")
	}
}
