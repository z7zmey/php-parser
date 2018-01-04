package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstList struct {
	position *node.Position
	Consts   []node.Node
}

func NewConstList(Consts []node.Node) node.Node {
	return &ConstList{
		nil,
		Consts,
	}
}

func (n ConstList) Attributes() map[string]interface{} {
	return nil
}

func (n ConstList) Position() *node.Position {
	return n.position
}

func (n ConstList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
