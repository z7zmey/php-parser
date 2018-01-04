package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassConstList struct {
	position  *node.Position
	Modifiers []node.Node
	Consts    []node.Node
}

func NewClassConstList(Modifiers []node.Node, Consts []node.Node) node.Node {
	return &ClassConstList{
		nil,
		Modifiers,
		Consts,
	}
}

func (n ClassConstList) Attributes() map[string]interface{} {
	return nil
}

func (n ClassConstList) Position() *node.Position {
	return n.position
}

func (n ClassConstList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClassConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			nn.Walk(vv)
		}
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
