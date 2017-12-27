package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func(n StmtList) Name() string {
	return "StmtList"
}

type StmtList struct {
	name  string
	stmts []node.Node
}

func NewStmtList(stmts []node.Node) node.Node {
	return StmtList{
		"StmtList",
		stmts,
	}
}

func (n StmtList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
