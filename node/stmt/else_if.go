package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ElseIf struct {
	position *node.Position
	Cond     node.Node
	Stmt     node.Node
}

func NewElseIf(Cond node.Node, Stmt node.Node) node.Node {
	return &ElseIf{
		nil,
		Cond,
		Stmt,
	}
}

func (n ElseIf) Attributes() map[string]interface{} {
	return nil
}

func (n ElseIf) Position() *node.Position {
	return n.position
}

func (n ElseIf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ElseIf) Walk(v node.Visitor) {
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
