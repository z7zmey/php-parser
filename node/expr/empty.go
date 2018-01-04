package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Empty struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewEmpty(Expression node.Node) node.Node {
	return &Empty{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n Empty) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Empty) Position() *node.Position {
	return n.position
}

func (n Empty) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Empty) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
