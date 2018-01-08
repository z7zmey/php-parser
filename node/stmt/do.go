package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Do struct {
	position *node.Position
	Stmt     node.Node
	Cond     node.Node
}

func NewDo(Stmt node.Node, Cond node.Node) *Do {
	return &Do{
		nil,
		Stmt,
		Cond,
	}
}

func (n *Do) Attributes() map[string]interface{} {
	return nil
}

func (n *Do) Position() *node.Position {
	return n.position
}

func (n *Do) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Do) Walk(v node.Visitor) {
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
