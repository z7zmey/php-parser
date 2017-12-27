package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Lnumber) Name() string {
	return "Lnumber"
}

type Lnumber struct {
	name  string
	token token.Token
}

func NewLnumber(token token.Token) node.Node {
	return Lnumber{
		"Lnumber",
		token,
	}
}

func (n Lnumber) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
