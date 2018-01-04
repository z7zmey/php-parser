package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type UseList struct {
	position *node.Position
	UseType  node.Node
	Uses     []node.Node
}

func NewUseList(UseType node.Node, Uses []node.Node) node.Node {
	return &UseList{
		nil,
		UseType,
		Uses,
	}
}

func (n UseList) Attributes() map[string]interface{} {
	return nil
}

func (n UseList) Position() *node.Position {
	return n.position
}

func (n UseList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n UseList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
