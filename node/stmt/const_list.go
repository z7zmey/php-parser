package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstList struct {
	attributes map[string]interface{}
	position   *node.Position
	consts     []node.Node
}

func NewConstList(consts []node.Node) node.Node {
	return &ConstList{
		map[string]interface{}{},
		nil,
		consts,
	}
}

func (n ConstList) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.consts != nil {
		vv := v.GetChildrenVisitor("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
