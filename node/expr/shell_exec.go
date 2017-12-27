package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n ShellExec) Name() string {
	return "ShellExec"
}

type ShellExec struct {
	name  string
	parts []node.Node
}

func NewShellExec(parts []node.Node) node.Node {
	return ShellExec{
		"ShellExec",
		parts,
	}
}

func (n ShellExec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.parts != nil {
		fmt.Fprintf(out, "\n%vparts:", indent+"  ")
		for _, nn := range n.parts {
			nn.Print(out, indent+"    ")
		}
	}
}
