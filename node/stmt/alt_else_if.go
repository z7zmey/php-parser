package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElseIf struct {
	Cond node.Node
	Stmt node.Node
}

func NewAltElseIf(Cond node.Node, Stmt node.Node) *AltElseIf {
	return &AltElseIf{
		Cond,
		Stmt,
	}
}

func (n *AltElseIf) Attributes() map[string]interface{} {
	return nil
}

func (n *AltElseIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
