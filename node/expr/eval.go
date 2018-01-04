package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewEval(Expression node.Node) node.Node {
	return &Eval{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n Eval) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Eval) Position() *node.Position {
	return n.position
}

func (n Eval) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Eval) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
