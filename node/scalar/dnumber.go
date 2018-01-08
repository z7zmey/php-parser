package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Dnumber struct {
	position *node.Position
	Value    string
}

func NewDnumber(Value string) *Dnumber {
	return &Dnumber{
		nil,
		Value,
	}
}

func (n *Dnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *Dnumber) Position() *node.Position {
	return n.position
}

func (n *Dnumber) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Dnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
