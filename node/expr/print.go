package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewPrint(expression node.Node) node.Node {
	return &Print{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n Print) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Print) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Print) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Print) Position() *node.Position {
	return n.position
}

func (n Print) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Print) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
