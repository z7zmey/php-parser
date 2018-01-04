package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostDec struct {
	position *node.Position
	Variable node.Node
}

func NewPostDec(Variable node.Node) node.Node {
	return &PostDec{
		nil,
		Variable,
	}
}

func (n PostDec) Attributes() map[string]interface{} {
	return nil
}

func (n PostDec) Position() *node.Position {
	return n.position
}

func (n PostDec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PostDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
