package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Else struct {
	attributes map[string]interface{}
	position   *node.Position
	stmt       node.Node
}

func NewElse(stmt node.Node) node.Node {
	return Else{
		map[string]interface{}{},
		nil,
		stmt,
	}
}

func (n Else) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Else) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Else) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Else) Position() *node.Position {
	return n.position
}

func (n Else) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Else) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
