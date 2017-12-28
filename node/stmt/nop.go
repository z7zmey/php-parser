package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Nop) Name() string {
	return "Nop"
}

type Nop struct {
	name  string
	token token.Token
}

func NewNop(token token.Token) node.Node {
	return Nop{
		"Nop",
		token,
	}
}

func (n Nop) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}

func (n Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	v.LeaveNode(n)
}
