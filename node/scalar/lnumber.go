package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Lnumber struct {
	attributes map[string]interface{}
	position   *node.Position
}

func NewLnumber(Value string) node.Node {
	return &Lnumber{
		map[string]interface{}{
			"Value": Value,
		},
		nil,
	}
}

func (n Lnumber) Attributes() map[string]interface{} {
	return n.attributes
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
