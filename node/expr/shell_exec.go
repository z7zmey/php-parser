package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ShellExec struct {
	node.SimpleNode
	parts []node.Node
}

func NewShellExec(parts []node.Node) node.Node {
	return ShellExec{
		node.SimpleNode{Name: "ShellExec", Attributes: make(map[string]string)},
		parts,
	}
}

func (n ShellExec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.parts != nil {
		fmt.Fprintf(out, "\n%vparts:", indent+"  ")
		for _, nn := range n.parts {
			nn.Print(out, indent+"    ")
		}
	}
}
