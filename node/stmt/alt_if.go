package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n AltIf) Name() string {
	return "AltIf"
}

type AltIf struct {
	name      string
	token     token.Token
	cond      node.Node
	stmt      node.Node
	elseAltIf []node.Node
	_else     node.Node
}

func NewAltIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltIf{
		"AltIf",
		token,
		cond,
		stmt,
		nil,
		nil,
	}
}

func (n AltIf) AddElseIf(elseAltIf node.Node) node.Node {
	if n.elseAltIf == nil {
		n.elseAltIf = make([]node.Node, 0)
	}

	n.elseAltIf = append(n.elseAltIf, elseAltIf)

	return n
}

func (n AltIf) SetElse(_else node.Node) node.Node {
	n._else = _else

	return n
}

func (n AltIf) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		n.cond.Print(out, indent+"    ")
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}

	if n.elseAltIf != nil {
		fmt.Fprintf(out, "\n%velseAltIfs:", indent+"  ")
		for _, nn := range n.elseAltIf {
			nn.Print(out, indent+"    ")
		}
	}
	if n._else != nil {
		fmt.Fprintf(out, "\n%velse:", indent+"  ")
		n._else.Print(out, indent+"    ")
	}
}
