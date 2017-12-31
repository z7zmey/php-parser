package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type UseList struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	useType    node.Node
	uses       []node.Node
}

func NewUseList(useType node.Node, uses []node.Node) node.Node {
	return UseList{
		"UseList",
		map[string]interface{}{},
		nil,
		useType,
		uses,
	}
}

func (n UseList) Name() string {
	return "UseList"
}

func (n UseList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n UseList) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n UseList) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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

	if n.useType != nil {
		vv := v.GetChildrenVisitor("useType")
		n.useType.Walk(vv)
	}

	if n.uses != nil {
		vv := v.GetChildrenVisitor("uses")
		for _, nn := range n.uses {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
