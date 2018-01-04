package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Interface struct {
	attributes    map[string]interface{}
	position      *node.Position
	InterfaceName node.Node
	Extends       []node.Node
	Stmts         []node.Node
}

func NewInterface(InterfaceName node.Node, Extends []node.Node, Stmts []node.Node, phpDocComment string) node.Node {
	return &Interface{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		InterfaceName,
		Extends,
		Stmts,
	}
}

func (n Interface) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Interface) Position() *node.Position {
	return n.position
}

func (n Interface) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Interface) Walk(v node.Visitor) {
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
