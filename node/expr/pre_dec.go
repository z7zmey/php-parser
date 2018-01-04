package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	attributes map[string]interface{}
	position   *node.Position
	Variable   node.Node
}

func NewPreDec(Variable node.Node) node.Node {
	return &PreDec{
		map[string]interface{}{},
		nil,
		Variable,
	}
}

func (n PreDec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PreDec) Position() *node.Position {
	return n.position
}

func (n PreDec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
