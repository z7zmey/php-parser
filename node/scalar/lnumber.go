package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Lnumber struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
}

func NewLnumber(value string) node.Node {
	return Lnumber{
		"Lnumber",
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n Lnumber) Name() string {
	return "Lnumber"
}

func (n Lnumber) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Lnumber) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Lnumber) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Lnumber) Position() *node.Position {
	return n.position
}

func (n Lnumber) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
