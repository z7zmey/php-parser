package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return ErrorSuppress{
		"ErrorSuppress",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n ErrorSuppress) Name() string {
	return "ErrorSuppress"
}

func (n ErrorSuppress) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ErrorSuppress) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ErrorSuppress) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ErrorSuppress) Position() *node.Position {
	return n.position
}

func (n ErrorSuppress) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ErrorSuppress) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
