package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n GroupUse) Name() string {
	return "GroupUse"
}

type GroupUse struct {
	name    string
	token   token.TokenInterface
	useType node.Node
	prefix  node.Node
	useList []node.Node
}

//TODO: stmts myst be []node.Node
func NewGroupUse(token token.TokenInterface, useType node.Node, prefix node.Node, useList []node.Node) node.Node {
	return GroupUse{
		"GroupUse",
		token,
		useType,
		prefix,
		useList,
	}
}

func (n GroupUse) SetToken(token token.TokenInterface) node.Node {
	n.token = token
	return n
}

func (n GroupUse) SetUseType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n GroupUse) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.GetStartLine(), n.token.GetEndLine(), n.token.GetValue())

	if n.useType != nil {
		fmt.Fprintf(out, "\n%vuse type:", indent+"  ")
		n.useType.Print(out, indent+"    ")
	}

	if n.prefix != nil {
		fmt.Fprintf(out, "\n%vprefix:", indent+"  ")
		n.prefix.Print(out, indent+"    ")
	}

	if n.useList != nil {
		fmt.Fprintf(out, "\n%vuse list:", indent+"  ")
		for _, nn := range n.useList {
			nn.Print(out, indent+"    ")
		}
	}
}
