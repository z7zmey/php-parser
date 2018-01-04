package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElse struct {
	attributes map[string]interface{}
	position   *node.Position
	Stmt       node.Node
}

func NewAltElse(Stmt node.Node) node.Node {
	return &AltElse{
		map[string]interface{}{},
		nil,
		Stmt,
	}
}

func (n AltElse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AltElse) Position() *node.Position {
	return n.position
}

func (n AltElse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
