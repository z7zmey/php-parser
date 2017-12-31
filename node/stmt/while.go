package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type While struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	token      token.Token
	cond       node.Node
	stmt       node.Node
}

func NewWhile(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return While{
		"While",
		map[string]interface{}{},
		nil,
		token,
		cond,
		stmt,
	}
}

func (n While) Name() string {
	return "While"
}

func (n While) Attributes() map[string]interface{} {
	return n.attributes
}

func (n While) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n While) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n While) Position() *node.Position {
	return n.position
}

func (n While) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n While) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
