package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Do struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	stmt       node.Node
	cond       node.Node
}

func NewDo(stmt node.Node, cond node.Node) node.Node {
	return Do{
		"Do",
		map[string]interface{}{},
		nil,
		stmt,
		cond,
	}
}

func (n Do) Name() string {
	return "Do"
}

func (n Do) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Do) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Do) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Do) Position() *node.Position {
	return n.position
}

func (n Do) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Do) Walk(v node.Visitor) {
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
