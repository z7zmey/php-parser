package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Constant struct {
	name         string
	attributes   map[string]interface{}
	position *node.Position
	constantName node.Node
	expr         node.Node
}

func NewConstant(constantName node.Node, expr node.Node) node.Node {
	return Constant{
		"Constant",
		map[string]interface{}{},
		nil,
		constantName,
		expr,
	}
}

func (n Constant) Name() string {
	return "Constant"
}

func (n Constant) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Constant) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Constant) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Constant) Position() *node.Position {
	return n.position
}

func (n Constant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Constant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constantName != nil {
		vv := v.GetChildrenVisitor("constantName")
		n.constantName.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
