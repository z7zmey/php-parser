package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Use struct {
	attributes map[string]interface{}
	position   *node.Position
	useType    node.Node
	use        node.Node
	alias      node.Node
}

func NewUse(useType node.Node, use node.Node, alias node.Node) node.Node {
	return Use{
		map[string]interface{}{},
		nil,
		useType,
		use,
		alias,
	}
}

func (n Use) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Use) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Use) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Use) Position() *node.Position {
	return n.position
}

func (n Use) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Use) SetUseType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n Use) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.GetChildrenVisitor("useType")
		n.useType.Walk(vv)
	}

	if n.use != nil {
		vv := v.GetChildrenVisitor("use")
		n.use.Walk(vv)
	}

	if n.alias != nil {
		vv := v.GetChildrenVisitor("alias")
		n.alias.Walk(vv)
	}

	v.LeaveNode(n)
}
