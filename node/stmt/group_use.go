package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type GroupUse struct {
	attributes map[string]interface{}
	position   *node.Position
	useType    node.Node
	prefix     node.Node
	useList    []node.Node
}

func NewGroupUse(useType node.Node, prefix node.Node, useList []node.Node) node.Node {
	return &GroupUse{
		map[string]interface{}{},
		nil,
		useType,
		prefix,
		useList,
	}
}

func (n GroupUse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n GroupUse) Position() *node.Position {
	return n.position
}

func (n GroupUse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n GroupUse) SetUseType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n GroupUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.GetChildrenVisitor("useType")
		n.useType.Walk(vv)
	}

	if n.prefix != nil {
		vv := v.GetChildrenVisitor("prefix")
		n.prefix.Walk(vv)
	}

	if n.useList != nil {
		vv := v.GetChildrenVisitor("useList")
		for _, nn := range n.useList {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
