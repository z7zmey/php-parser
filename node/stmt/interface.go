package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Interface struct {
	attributes    map[string]interface{}
	position      *node.Position
	interfaceName node.Node
	extends       []node.Node
	stmts         []node.Node
}

func NewInterface(interfaceName node.Node, extends []node.Node, stmts []node.Node, phpDocComment string) node.Node {
	return &Interface{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		interfaceName,
		extends,
		stmts,
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

	if n.interfaceName != nil {
		vv := v.GetChildrenVisitor("interfaceName")
		n.interfaceName.Walk(vv)
	}

	if n.extends != nil {
		vv := v.GetChildrenVisitor("extends")
		for _, nn := range n.extends {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
