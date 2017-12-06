package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Class struct {
	node.SimpleNode
	token      token.Token
	modifiers  []string
	args       []node.Node
	extends    node.Node
	implements []node.Node
	stmts      []node.Node
}

//TODO: stmts myst be []node.Node
func NewClass(token token.Token, modifiers []string, args []node.Node, extends node.Node, implements []node.Node, stmts []node.Node) node.Node {
	return Class{
		node.SimpleNode{Name: "Class", Attributes: make(map[string]string)},
		token,
		modifiers,
		args,
		extends,
		implements,
		stmts,
	}
}

func (n Class) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.modifiers != nil {
		fmt.Fprintf(out, "\n%vmotifiers:", indent+"  ")
		for _, nn := range n.modifiers {
			fmt.Fprintf(out, "\n%v%q", indent+"    ", nn)
		}
	}

	if n.args != nil {
		fmt.Fprintf(out, "\n%vargs:", indent+"  ")
		for _, nn := range n.args {
			nn.Print(out, indent+"    ")
		}
	}

	if n.extends != nil {
		fmt.Fprintf(out, "\n%vextends:", indent+"  ")
		n.extends.Print(out, indent+"    ")
	}

	if n.implements != nil {
		fmt.Fprintf(out, "\n%vimplements:", indent+"  ")
		for _, nn := range n.implements {
			nn.Print(out, indent+"    ")
		}
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
