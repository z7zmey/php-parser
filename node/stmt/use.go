package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Use struct {
	node.SimpleNode
	useType node.Node
	name    node.Node
	alias   token.TokenInterface
}

func NewUse(useType node.Node, name node.Node, alias token.TokenInterface) node.Node {
	return Use{
		node.SimpleNode{Name: "Use", Attributes: make(map[string]string)},
		useType,
		name,
		alias,
	}
}

func (n Use) SetType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n Use) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.useType != nil {
		fmt.Fprintf(out, "\n%vtype:", indent+"  ")
		n.useType.Print(out, indent+"    ")
	}

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}

	if n.alias != nil {
		fmt.Fprintf(out, "\n%valias: %q", indent+"  ", n.alias.GetValue())
	}
}
