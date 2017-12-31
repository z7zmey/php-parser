package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Dnumber struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
}

func NewDnumber(value string) node.Node {
	return Dnumber{
		"Dnumber",
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n Dnumber) Name() string {
	return "Dnumber"
}

func (n Dnumber) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Dnumber) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Dnumber) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Dnumber) Position() *node.Position {
	return n.position
}

func (n Dnumber) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Dnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
