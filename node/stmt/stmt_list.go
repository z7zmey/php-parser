package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StmtList struct {
	Stmts []node.Node
}

func NewStmtList(Stmts []node.Node) *StmtList {
	return &StmtList{
		Stmts,
	}
}

func (n *StmtList) Attributes() map[string]interface{} {
	return nil
}

func (n *StmtList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
