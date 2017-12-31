package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Foreach struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	expr       node.Node
	key        node.Node
	variable   node.Node
	stmt       node.Node
}

func NewForeach(expr node.Node, key node.Node, variable node.Node, stmt node.Node, byRef bool) node.Node {
	return Foreach{
		"Foreach",
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		expr,
		key,
		variable,
		stmt,
	}
}

func (n Foreach) Name() string {
	return "Foreach"
}

func (n Foreach) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Foreach) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Foreach) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Foreach) Position() *node.Position {
	return n.position
}

func (n Foreach) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Foreach) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.key != nil {
		vv := v.GetChildrenVisitor("key")
		n.key.Walk(vv)
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
