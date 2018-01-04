package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type For struct {
	attributes map[string]interface{}
	position   *node.Position
	Init       []node.Node
	Cond       []node.Node
	Loop       []node.Node
	Stmt       node.Node
}

func NewFor(Init []node.Node, Cond []node.Node, Loop []node.Node, Stmt node.Node) node.Node {
	return &For{
		map[string]interface{}{},
		nil,
		Init,
		Cond,
		Loop,
		Stmt,
	}
}

func (n For) Attributes() map[string]interface{} {
	return n.attributes
}

func (n For) Position() *node.Position {
	return n.position
}

func (n For) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n For) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Init != nil {
		vv := v.GetChildrenVisitor("Init")
		for _, nn := range n.Init {
			nn.Walk(vv)
		}
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		for _, nn := range n.Cond {
			nn.Walk(vv)
		}
	}

	if n.Loop != nil {
		vv := v.GetChildrenVisitor("Loop")
		for _, nn := range n.Loop {
			nn.Walk(vv)
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
