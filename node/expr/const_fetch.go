package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	position *node.Position
	Constant node.Node
}

func NewConstFetch(Constant node.Node) node.Node {
	return &ConstFetch{
		nil,
		Constant,
	}
}

func (n ConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n ConstFetch) Position() *node.Position {
	return n.position
}

func (n ConstFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Constant != nil {
		vv := v.GetChildrenVisitor("Constant")
		n.Constant.Walk(vv)
	}

	v.LeaveNode(n)
}
