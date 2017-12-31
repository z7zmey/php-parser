package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElse struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	stmt       node.Node
}

func NewAltElse(stmt node.Node) node.Node {
	return AltElse{
		"AltElse",
		map[string]interface{}{},
		nil,
		stmt,
	}
}

func (n AltElse) Name() string {
	return "AltElse"
}

func (n AltElse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AltElse) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n AltElse) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n AltElse) Position() *node.Position {
	return n.position
}

func (n AltElse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
