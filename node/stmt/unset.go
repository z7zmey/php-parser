package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Unset struct {
	position *node.Position
	Vars     []node.Node
}

func NewUnset(Vars []node.Node) *Unset {
	return &Unset{
		nil,
		Vars,
	}
}

func (n *Unset) Attributes() map[string]interface{} {
	return nil
}

func (n *Unset) Position() *node.Position {
	return n.position
}

func (n *Unset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Unset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		vv := v.GetChildrenVisitor("Vars")
		for _, nn := range n.Vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
