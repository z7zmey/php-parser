package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Interface struct {
	position      *node.Position
	PhpDocComment string
	InterfaceName node.Node
	Extends       []node.Node
	Stmts         []node.Node
}

func NewInterface(InterfaceName node.Node, Extends []node.Node, Stmts []node.Node, PhpDocComment string) *Interface {
	return &Interface{
		nil,
		PhpDocComment,
		InterfaceName,
		Extends,
		Stmts,
	}
}

func (n *Interface) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n *Interface) Position() *node.Position {
	return n.position
}

func (n *Interface) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Interface) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InterfaceName != nil {
		vv := v.GetChildrenVisitor("InterfaceName")
		n.InterfaceName.Walk(vv)
	}

	if n.Extends != nil {
		vv := v.GetChildrenVisitor("Extends")
		for _, nn := range n.Extends {
			nn.Walk(vv)
		}
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
