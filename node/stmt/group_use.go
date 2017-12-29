package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type GroupUse struct {
	name       string
	attributes map[string]interface{}
	useType    node.Node
	prefix     node.Node
	useList    []node.Node
}

func NewGroupUse(useType node.Node, prefix node.Node, useList []node.Node) node.Node {
	return GroupUse{
		"GroupUse",
		map[string]interface{}{},
		useType,
		prefix,
		useList,
	}
}

func (n GroupUse) Name() string {
	return "GroupUse"
}

func (n GroupUse) Attributes() map[string]interface{} {
	return n.attributes
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
