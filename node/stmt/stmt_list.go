package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type StmtList struct {
	node.SimpleNode
	stmts []node.Node
}

func NewStmtList(stmts []node.Node) node.Node {
	return StmtList{
		node.SimpleNode{Name: "StmtList", Attributes: make(map[string]string)},
		stmts,
	}
}

func (n StmtList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
