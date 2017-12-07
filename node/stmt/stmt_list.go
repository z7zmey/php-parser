package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type StmtList struct {
	node.SimpleNode
	openBracket  token.Token
	closeBracket token.Token
	stmts        []node.Node
}

func NewStmtList(openBracket token.Token, closeBracket token.Token, stmts []node.Node) node.Node {
	return StmtList{
		node.SimpleNode{Name: "StmtList", Attributes: make(map[string]string)},
		openBracket,
		closeBracket,
		stmts,
	}
}

func (n StmtList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d]", indent, n.Name, n.openBracket.StartLine, n.closeBracket.EndLine)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
