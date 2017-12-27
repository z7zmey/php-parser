package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Use) Name() string {
	return "Use"
}

type Use struct {
	name    string
	useType node.Node
	use     node.Node
	alias   token.TokenInterface
}

func NewUse(useType node.Node, use node.Node, alias token.TokenInterface) node.Node {
	return Use{
		"Use",
		useType,
		use,
		alias,
	}
}

func (n Use) SetType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n Use) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.useType != nil {
		fmt.Fprintf(out, "\n%vtype:", indent+"  ")
		n.useType.Print(out, indent+"    ")
	}

	if n.use != nil {
		fmt.Fprintf(out, "\n%vuse:", indent+"  ")
		n.use.Print(out, indent+"    ")
	}

	if n.alias != nil {
		fmt.Fprintf(out, "\n%valias: %q", indent+"  ", n.alias.GetValue())
	}
}
