package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	varName    node.Node
}

func NewVariable(varName node.Node) node.Node {
	return Variable{
		"Variable",
		map[string]interface{}{},
		nil,
		varName,
	}
}

func (n Variable) Name() string {
	return "Variable"
}

func (n Variable) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Variable) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Variable) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Variable) Position() *node.Position {
	return n.position
}

func (n Variable) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.varName != nil {
		vv := v.GetChildrenVisitor("varName")
		n.varName.Walk(vv)
	}

	v.LeaveNode(n)
}
