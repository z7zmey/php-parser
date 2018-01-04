package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ClusureUse struct {
	attributes map[string]interface{}
	position   *node.Position
	Variable   node.Node
}

func NewClusureUse(Variable node.Node, byRef bool) node.Node {
	return &ClusureUse{
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		Variable,
	}
}

func (n ClusureUse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClusureUse) Position() *node.Position {
	return n.position
}

func (n ClusureUse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClusureUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
