package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n EncapsedStringPart) Name() string {
	return "EncapsedStringPart"
}

type EncapsedStringPart struct {
	name  string
	token token.Token
}

func NewEncapsedStringPart(t token.Token) node.Node {
	return EncapsedStringPart{
		"EncapsedStringPart",
		t,
	}
}

func (n EncapsedStringPart) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
