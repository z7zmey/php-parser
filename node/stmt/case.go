package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Case struct {
	Cond  node.Node
	Stmts []node.Node
}

func NewCase(Cond node.Node, Stmts []node.Node) *Case {
	return &Case{
		Cond,
		Stmts,
	}
}

func (n *Case) Attributes() map[string]interface{} {
	return nil
}

func (n *Case) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
