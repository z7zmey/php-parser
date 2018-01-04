package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Continue struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewContinue(expr node.Node) node.Node {
	return &Continue{
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Continue) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Continue) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Continue) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Continue) Position() *node.Position {
	return n.position
}

func (n Continue) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Continue) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
